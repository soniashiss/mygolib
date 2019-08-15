package requests

import (
	"net/http"
	"io/ioutil"
)

func SendGET(uri string, param string) ([]byte, error) {
	newurl := uri + param

	client := &http.Client{
		Timeout: requestTimeout,
	}
	resp, err := client.Get(newurl)
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

func SendGETWithHeader(uri string, param string, header map[string]string) ([]byte, error) {
	newurl := uri + param

	client := &http.Client{
		Timeout: requestTimeout,
	}

	req, _ := http.NewRequest("GET", newurl, nil)

	for k, v := range header {
		req.Header.Add(k,v)
	}

	resp, err := client.Do(req)
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

func SendGETWithHeaderCodeReturn(uri string, param string, header map[string]string) (int, []byte, error) {
	newurl := uri + param

	client := &http.Client{
		Timeout: requestTimeout,
	}

	req, _ := http.NewRequest("GET", newurl, nil)

	for k, v := range header {
		req.Header.Add(k,v)
	}

	resp, err := client.Do(req)
	code := int(404)
	if resp != nil {
		code = resp.StatusCode
		defer resp.Body.Close()
	}
	if err != nil {
		return code, []byte{}, err
	} else {
		result, err := ioutil.ReadAll(resp.Body)
		//fmt.Println(string(result))
		if err != nil {
			return code, []byte{}, err
		} else {
			return code, result, nil
		}
	}
}
