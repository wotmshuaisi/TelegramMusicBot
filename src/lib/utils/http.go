package utils

import (
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	client = &http.Client{
		Timeout: time.Second * 30,
	}
)

// HTTPGetJSON http request for json data
func HTTPGetJSON(url string) ([]byte, error) {
	var jsonData []byte
	resp, err := client.Do(newRequest(http.MethodGet, url, nil))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if jsonData, err = ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	}

	return jsonData, nil
}

// HTTPPostJSON http request for json data
func HTTPPostJSON(url string, data io.Reader) ([]byte, error) {
	var jsonData []byte
	resp, err := client.Do(newRequest(http.MethodPost, url, data))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if jsonData, err = ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	}

	return jsonData, nil
}

func newRequest(method, url string, data io.Reader) *http.Request {
	req, _ := http.NewRequest(method, url, data)
	req.Header.Set("user-agen", "Mozilla/5.0 (Windows NT 10.0; Win64; rv:64.0) Gecko/20100101 Firefox/64.0")
	return req
}
