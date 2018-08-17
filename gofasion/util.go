package gofasion

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

func HttpGet(targetUrl string, params url.Values) (bs []byte, err error) {
	resp, err := http.Get(targetUrl + "?" + params.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	ret, err := ioutil.ReadAll(resp.Body)
	return ret, err
}
