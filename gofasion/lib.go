package gofasion

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"strconv"
	"time"
	"unicode"
)

func (self *Fasion) generateMd5() string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(strconv.FormatInt(time.Now().UnixNano(), 10)))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func (self *Fasion) parseJson() (map[string]interface{}, error) {
	if !json.Valid([]byte(self.rawJson)) || self.isValidNum(self.rawJson) {
		ret := make(map[string]interface{}, 0)
		self.currentKey = self.generateMd5()
		ret[self.currentKey] = self.rawJson
		return ret, nil
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

func (self *Fasion) isValidNum(str string) bool {
	if str == "" {
		return false
	}
	hasBegin := false
	for i, ch := range str {
		if i == 0 && (ch != '-' && !unicode.IsDigit(ch)) {
			return false
		}
		if i == 0 && ch == '-' {
			continue
		}
		if i == 0 && ch == '.' {
			return false
		}
		if ch == '.' && hasBegin {
			return false
		}
		if ch == '.' && !hasBegin {
			hasBegin = true
			continue
		}
		if !unicode.IsDigit(ch) {
			return false
		}
	}
	return true
}
