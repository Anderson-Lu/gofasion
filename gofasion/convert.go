package gofasion

import (
	"encoding/json"
	"strconv"
	"strings"
)

func (self *Fasion) initCur() {
	if self.current == nil {
		curNode, err := self.parseJson()
		self.current = curNode
		self.errInfo = err
	}
}

func (self *Fasion) initArray() {
	if self.current == nil {
		curNode, err := self.parseArray()
		self.current = curNode
		self.errInfo = err
	}
}

func (self *Fasion) ValueStr() string {
	self.initCur()
	if val, ok := self.current.(string); ok {
		return strings.Replace(val, "\"", "", -1)
	}
	return strings.Replace(self.rawJson, "\"", "", -1)
}

func (self *Fasion) ValueInt64() int64 {
	self.initCur()
	if val, ok := self.current.(int64); ok {
		return val
	}
	if n, ok := strconv.ParseInt(self.rawJson, 10, 64); ok == nil {
		return n
	}
	return 0
}

func (self *Fasion) ValueInt32() int32 {
	self.initCur()
	if val, ok := self.current.(int32); ok {
		return val
	}
	if n, ok := strconv.ParseInt(self.rawJson, 10, 32); ok == nil {
		return int32(n)
	}
	return 0
}

func (self *Fasion) ValueInt() int {
	self.initCur()
	if val, ok := self.current.(int); ok {
		return val
	}
	if n, ok := strconv.Atoi(self.rawJson); ok == nil {
		return n
	}
	return 0
}

func (self *Fasion) Json() string {
	return strings.Trim(self.rawJson, " ")
}

func (self *Fasion) Array() []*Fasion {
	self.initArray()
	result := make([]*Fasion, 0)
	if val, ok := self.current.([]interface{}); ok {
		for _, v := range val {
			rawJson, err := self.toJson(v)
			if err == nil {
				result = append(result, NewFasion(rawJson))
			}
		}
	}
	return result
}

func (self *Fasion) ValueBool() bool {
	self.initCur()
	valueStr := self.ValueStr()
	if valueStr == "" {
		return false
	}
	if strings.ToLower(valueStr) == "false" {
		return false
	}
	if strings.ToLower(valueStr) == "true" {
		return true
	}
	return false
}

func (self *Fasion) Value(dest interface{}) error {
	return json.Unmarshal([]byte(self.rawJson), &dest)
}
