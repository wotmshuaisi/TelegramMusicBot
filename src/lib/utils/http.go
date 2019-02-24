package utils

import (
	"io/ioutil"
	"net/http"
	"time"
)

var (
	client = &http.Client{
		Timeout: time.Second * 5,
	}
)

// GetJSON http request for json data
func GetJSON(url string) ([]byte, error) {
	var jsonData []byte
	resp, err := client.Do(newGetRequest(url))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if jsonData, err = ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	}

	return jsonData, nil
}

func newGetRequest(url string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("user-agen", "Mozilla/5.0 (Windows NT 10.0; Win64; rv:64.0) Gecko/20100101 Firefox/64.0")
	return req
}
