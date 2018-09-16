package gofasion

import (
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"strings"
)

var encodingPlacer = strings.NewReplacer("\"", "", "\\u0026", "&", "\\u003c", "<", "\\u003e", ">", "\\u003d", "=")

func httpGet(targetUrl string, params url.Values) (bs []byte, err error) {
	resp, err := http.Get(targetUrl + "?" + params.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	ret, err := ioutil.ReadAll(resp.Body)
	return ret, err
}

func encodingParser(raw string) string {
	return encodingPlacer.Replace(raw)
}

func round(raw float64, spec int) float64 {
	n10 := math.Pow10(spec)
	return math.Trunc((raw+0.5/n10)*n10) / n10
}
