package gofasion

import (
	"encoding/json"
	"fmt"
)

func (self *Fasion) parseJson() (map[string]interface{}, error) {
	if !json.Valid([]byte(self.rawJson)) {
		return nil, fmt.Errorf("invalid json format")
	}
	var result map[string]interface{}
	e := _unmarshalFunc([]byte(self.rawJson), &result)
	return result, e
}

func (self *Fasion) parseArray() ([]interface{}, error) {
	var result []interface{}
	e := _unmarshalFunc([]byte(self.rawJson), &result)
	return result, e
}

func (self *Fasion) toJson(target interface{}) (string, error) {
	if val, ok := target.(string); ok {
		return val, nil
	}
	bs, err := _marshalFunc(target)
	return string(bs), err
}
