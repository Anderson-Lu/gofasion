package gofasion

import (
	"encoding/json"
	"net/url"
	"strings"
)

var _unmarshalFunc func([]byte, interface{}) error
var _marshalFunc func(interface{}) ([]byte, error)

type Fasion struct {
	rawJson    string
	errInfo    error
	current    interface{}
	currentKey string
	exists     bool
}

//Setup your custom marshal and unmarshal methods
func SetJsonParser(customMarshal func(interface{}) ([]byte, error), customUnmarshal func([]byte, interface{}) error) {
	_unmarshalFunc = customUnmarshal
	_marshalFunc = customMarshal
}

func NewFasion(rawJson string) *Fasion {
	_unmarshalFunc = json.Unmarshal
	_marshalFunc = json.Marshal
	return &Fasion{
		rawJson: rawJson,
	}
}

func newFasion(rawJson string, exists bool) *Fasion {
	_unmarshalFunc = json.Unmarshal
	_marshalFunc = json.Marshal
	return &Fasion{
		rawJson: rawJson,
		exists:  exists,
	}
}

func NewFasionFromBytes(rawJson []byte) *Fasion {
	return &Fasion{
		rawJson: string(rawJson),
	}
}

func NewFasionFromUrl(targetUrl string, params url.Values) *Fasion {
	if params == nil {
		params = url.Values{}
	}
	bs, err := httpGet(targetUrl, params)
	if err != nil {
		return &Fasion{
			rawJson: "",
			errInfo: err,
		}
	}
	return &Fasion{
		rawJson: string(bs),
	}
}

//Get specific node
func (self *Fasion) Get(key string) *Fasion {
	curMap, err := self.parseJson()
	if err != nil {
		self.errInfo = err
		self.current = nil
	}
	if v, ok := curMap[key]; ok {
		rawJson, err := self.toJson(v)
		if err == nil {
			return newFasion(rawJson, true)
		}
	}
	return newFasion("", false)
}

//Get node directly via absolute path like "node1.node2.node3"
func (self *Fasion) GetFromPath(dir string) *Fasion {
	paths := strings.Split(dir, ".")
	var ret *Fasion
	for _, path := range paths {
		if ret == nil {
			ret = self.Get(path)
		} else {
			ret = ret.Get(path)
		}
	}
	return ret
}

//Judge whether the JSON format is correct.
func (self *Fasion) IsValidJson() bool {
	return json.Valid([]byte(self.rawJson))
}
