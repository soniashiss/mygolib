package requests

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func SendPOST(url string, header map[string]string, param []byte) (int, string, error) {
	body := bytes.NewBuffer(param)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return 500, "", err
	}

	for k, v := range header {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return 500, "", err
	} else {
		result, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return 500, "", err
		} else {
			return resp.StatusCode, string(result), nil
		}
	}
}
