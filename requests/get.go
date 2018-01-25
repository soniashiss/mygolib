package requests

import (
	"net/http"
	"io/ioutil"
)

func SendGET(uri string, param string) ([]byte, error) {
	newurl := uri + param

	resp, err := http.Get(newurl)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return []byte{}, err
	} else {
		result, err := ioutil.ReadAll(resp.Body)
		//fmt.Println(string(result))
		if err != nil {
			return []byte{}, err
		} else {
			return result, nil
		}
	}
}
