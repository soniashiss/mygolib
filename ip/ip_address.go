package ip

import (
	"errors"
	"net"
)

func GetLocalIP4() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	//fmt.Println(addrs)
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback(){
			//为nil说明不是ipv4地址
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", errors.New("cannot find local ip")
}
