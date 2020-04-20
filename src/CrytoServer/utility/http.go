package utility

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

func HttpReq(method string, url string, body []byte) ([]byte, error) {

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("Error in http request for vm passivation or activation:", err)
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("failed to execute kubectl api request ", err)
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)
}
