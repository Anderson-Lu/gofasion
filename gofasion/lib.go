package gofasion

import (
	"encoding/json"
)

func (self *Fasion) parseJson() (map[string]interface{}, error) {
	var result map[string]interface{}
	e := json.Unmarshal([]byte(self.rawJson), &result)
	return result, e
}

func (self *Fasion) parseArray() ([]interface{}, error) {
	var result []interface{}
	e := json.Unmarshal([]byte(self.rawJson), &result)
	return result, e
}

func (self *Fasion) toJson(target interface{}) (string, error) {
	if val, ok := target.(string); ok {
		return val, nil
	}
	bs, err := json.Marshal(target)
	return string(bs), err
}
