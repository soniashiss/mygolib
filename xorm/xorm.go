package orm

import (
	"io"
	"time"

	//import mysql, do not need init
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

var (
	Orm  *localOrm
	Stop int
)

func InitiateOrm(username, password, tcpAddr, databaseName string) {

	url := username + ":" + password + "@tcp(" + tcpAddr + ")/" + databaseName + "?charset=utf8"
	_xorm, _ := xorm.NewEngine("mysql", url)
	Orm = &localOrm{_xorm, 4}

	Orm.ShowSQL(true)
	Orm.SetLogLevel(core.LOG_ERR)
	Orm.SetMaxIdleConns(10)
	Orm.SetMaxOpenConns(20)
	Orm.SetConnMaxLifetime(4 * time.Hour)
	//default logger is directed to os.stdOut
	go keepAlive(&Stop)
}
func SetLogger(out io.Writer, level core.LogLevel) {
	logger := xorm.NewSimpleLogger(out)
	logger.SetLevel(level)
	Orm.SetLogger(logger)
	Orm.Logger().Infof("set logger successfully\n")
}

type localOrm struct {
	*xorm.Engine
	RETRY int
}

func (self *localOrm) SetRetryTime(retryTime int) {
	self.RETRY = retryTime
}

func (self *localOrm) Find(beans interface{}, condiBeans ...interface{}) error {
	var ans error
	if ans = self.Engine.Find(beans, condiBeans); ans != nil {
		for i := 0; i < self.RETRY; i++ {
			if ans = self.Engine.Find(beans, condiBeans); ans == nil {
				break
			}
		}
	}
	return ans
}

func (self *localOrm) FindAndCount(rowsSlicePtr interface{}, condiBean ...interface{}) (int64, error) {
	var err error
	var count int64
	count, err = self.Engine.FindAndCount(rowsSlicePtr, condiBean)
	if err != nil {
		for i := 0; i < self.RETRY; i++ {
			count, err = self.Engine.FindAndCount(rowsSlicePtr, condiBean)
			if err == nil {
				break
			}
		}
	}
	return count, err
}
func (self *localOrm) Get(bean interface{}) (bool, error) {
	var ans1 bool
	var ans2 error
	if ans1, ans2 = self.Engine.Get(bean); ans2 != nil {
		for i := 0; i < self.RETRY; i++ {
			if ans1, ans2 = self.Engine.Get(bean); ans2 == nil {
				break
			}
		}
	}
	return ans1, ans2
}

func (self *localOrm) Update(bean interface{}, condiBeans ...interface{}) (int64, error) {
	var ans1 int64
	var ans2 error
	if ans1, ans2 = self.Engine.Update(bean, condiBeans); ans2 != nil {
		for i := 0; i < self.RETRY; i++ {
			if ans1, ans2 = self.Engine.Update(bean, condiBeans); ans2 == nil {
				break
			}
		}
	}
	return ans1, ans2
}

func (self *localOrm) Count(bean ...interface{}) (int64, error) {
	var ans1 int64
	var ans2 error
	if ans1, ans2 = self.Engine.Count(bean); ans2 != nil {
		for i := 0; i < self.RETRY; i++ {
			if ans1, ans2 = self.Engine.Count(bean); ans2 == nil {
				break
			}
		}
	}
	return ans1, ans2
}

func (self *localOrm) Exist(bean ...interface{}) (bool, error) {
	var has bool
	var err error
	if has, err = self.Engine.Exist(bean); err != nil {
		for i := 0; i < self.RETRY; i++ {
			if has, err = self.Engine.Exist(bean); err == nil {
				break
			}
		}
	}
	return has, err
}

//func (self *localOrm)Insert(beans ...interface{}) (int64, error) {
//	var affected int64
//	var err error
//	for i := 0; i <= self.RETRY; i++ {
//		if affected, err = self.Engine.Insert(beans); err != nil {
//			break
//		}
//	}
//	return affected, err
//}

func keepAlive(sp *int) {
	//defer Orm.Close()
	for *sp != 1 {
		time.Sleep(10 * time.Second)
		if err := Orm.Ping(); err != nil {
			Orm.Logger().Warnf("Ping err: %v", err.Error())
		}
	}
}
