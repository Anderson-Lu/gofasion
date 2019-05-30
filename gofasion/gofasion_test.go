package gofasion

import (
	"fmt"
	"testing"
)

var (
	rawJson string
	root    *Fasion
)

func init() {
	rawJson = `
		{
			"k1":-1,
			"bool":true,
			"bool1":false,
			"bool2":"true",
			"bool3":"false",
			"str1":"",
			"str2":"str1",
			"array1":[1,2,3,4],
			"float1": 1.1,
			"float2": 1.1155
		}
	`
	root = NewFasion(rawJson)
}

func TestBase(t *testing.T) {
	if root.Get("k1").ValueInt64() != -1 {
		t.Errorf("test failed")
	}
	if root.Get("k1").ValueInt32() != -1 {
		t.Errorf("test failed")
	}
	if root.Get("k1").ValueInt16() != -1 {
		t.Errorf("test failed")
	}
	if root.Get("k1").ValueInt() != -1 {
		t.Errorf("test failed")
	}
	if root.Get("float1").ValueFloat64() != 1.1 {
		t.Errorf("test failed")
	}
	if root.Get("float1").ValueFloat32() != 1.1 {
		t.Errorf("test failed")
	}
	if root.Get("float2").ValueFloat32N(3) != 1.115 {
		t.Errorf("test failed")
	}
	if root.Get("float2").ValueFloat32N(0) != 1.1155 {
		t.Errorf("test failed")
	}
	if root.Get("float2").ValueFloat64N(2) != 1.12 {
		t.Errorf("test failed")
	}
	if root.Get("float2").ValueFloat64N(-1) != 1.1155 {
		t.Errorf("test failed")
	}
	if root.Get("array1").Json() != `[1,2,3,4]` {
		t.Errorf("test failed")
	}
	if root.Get("err").ValueDefaultInt(1) != 1 {
		t.Errorf("test failed")
	}
	if root.Get("err").ValueDefaultInt16(1) != 1 {
		t.Errorf("test failed")
	}
	if root.Get("err").ValueDefaultInt32(1) != 1 {
		t.Errorf("test failed")
	}
	if root.Get("err").ValueDefaultInt64(1) != 1 {
		t.Errorf("test failed")
	}
	if root.Get("k1").ValueDefaultInt(1) != -1 {
		t.Errorf("test failed")
	}
	if root.Get("k1").ValueDefaultInt16(1) != -1 {
		t.Errorf("test failed")
	}
	if root.Get("k1").ValueDefaultInt32(1) != -1 {
		t.Errorf("test failed")
	}
	if root.Get("k1").ValueDefaultInt64(1) != -1 {
		t.Errorf("test failed")
	}
	if root.Get("err").ValueDefaultFloat32(1.1) != 1.1 {
		t.Errorf("test failed")
	}
	if root.Get("err").ValueDefaultFloat64(1.1) != 1.1 {
		t.Errorf("test failed")
	}
	if root.Get("float2").ValueDefaultFloat32(1.1) != 1.1155 {
		t.Errorf("test failed")
	}
	if root.Get("float2").ValueDefaultFloat64(1.1) != 1.1155 {
		t.Errorf("test failed")
	}
	if root.Get("err").ValueDefaultBool(true) != true {
		t.Errorf("test failed")
	}
	if root.Get("bool").ValueDefaultBool(false) != true {
		t.Errorf("test failed")
	}
	root.ArrayForEach(func(idx int, fsion *Fasion) {
		if idx+1 != fsion.ValueInt() {
			t.Errorf("test failed")
		}
	})
	if ok, val := root.Get("str2").ValStr(); !ok || val != "str1" {
		t.Errorf("test failed,%s,%v", val, ok)
	}
	if ok, val := root.Get("none").ValStr(); ok || val != "" {
		t.Errorf("test failed,%s,%v", val, ok)
	}
	if ok, val := root.Get("str1").ValStr(); !ok || val != "" {
		t.Errorf("test failed,%s,%v", val, ok)
	}
	if ok, val := root.Get("none").ValInt64(); ok || val != 0 {
		t.Errorf("test failed,%d,%v", val, ok)
	}
	if ok, val := root.Get("none").ValInt32(); ok || val != 0 {
		t.Errorf("test failed,%d,%v", val, ok)
	}
	if ok, val := root.Get("none").ValInt16(); ok || val != 0 {
		t.Errorf("test failed,%d,%v", val, ok)
	}
	if ok, val := root.Get("none").ValInt(); ok || val != 0 {
		t.Errorf("test failed,%d,%v", val, ok)
	}
	if ok, val := root.Get("none").ValFloat64(); ok || val != 0 {
		t.Errorf("test failed,%f,%v", val, ok)
	}
	if ok, val := root.Get("none").ValFloat64N(2); ok || val != 0 {
		t.Errorf("test failed,%f,%v", val, ok)
	}
	if ok, val := root.Get("none").ValFloat32N(2); ok || val != 0 {
		t.Errorf("test failed,%f,%v", val, ok)
	}
	if ok, val := root.Get("none").ValFloat32(); ok || val != 0 {
		t.Errorf("test failed,%f,%v", val, ok)
	}
	if ok, val := root.Get("none").ValBool(); ok || val {
		t.Errorf("test failed,%v,%v", val, ok)
	}
}

func TestIsValidJson(t *testing.T) {
	if !root.IsValidJson() {
		t.Errorf("failed to test isvalid json")
	}
	if root.Get("xxx").IsValidJson() {
		t.Errorf("failed to test isvalid json")
	}
	if root.Get("str1").IsValidJson() {
		t.Errorf("failed to test isvalid json")
	}
	if !root.Get("array1").IsValidJson() {
		t.Errorf("failed to test isvalid json")
	}
	if !root.Get("k1").IsValidJson() {
		t.Errorf("failed to test isvalid json")
	}
}

func TestValueArray(t *testing.T) {
	arr := root.Get("array1").Array()
	if len(arr) != 4 {
		t.Errorf("failed to test value array")
	}
	for i, v := range arr {
		if v.ValueInt() != i+1 {
			t.Errorf("failed to test value array,%d", v.ValueInt())
		}
	}
}

func TestValueStr(t *testing.T) {
	if root.Get("str1").ValueStr() != "" || root.Get("str1").ValueDefaultStr("k") != "k" {
		t.Errorf("Test get value str failed")
	}
	if root.Get("str2").ValueDefaultStr("k11") == "k11" {
		t.Errorf("Test get keys failed")
	}
}

func TestGetKeys(t *testing.T) {
	keys := root.Keys()
	if len(keys) != 10 {
		t.Errorf("Test get keys failed")
	}
	if len(root.Get("k1").Keys()) != 0 {
		t.Errorf("Test get keys failed")
	}
	if !root.HasKey("k1") {
		t.Errorf("Test get keys failed")
	}
	if root.HasKey("k1errr") {
		t.Errorf("Test get keys failed")
	}
	if len(root.Get("k1").Keys()) != 0 {
		t.Errorf("Test get keys failed")
	}
}

func TestValue(t *testing.T) {
	var ret struct {
		K1    int    `json:"k1"`
		Bool  bool   `json:"bool"`
		Bool1 bool   `json:"bool1"`
		Bool2 string `json:"bool2"`
		Bool3 string `json:"bool3"`
		Str1  string `json:"str1"`
	}
	err := root.Value(&ret)
	if err != nil {
		t.Errorf("Test value failed, %+v", err)
	}
	if ret.K1 != -1 || !ret.Bool || ret.Bool1 || !(ret.Bool2 == "true") || !(ret.Bool3 == "false") || ret.Str1 != "" {
		t.Errorf("Test value failed")
	}
}

func TestGet(t *testing.T) {
	fsion := root.Get("k1")
	fmt.Println("===", fsion.rawJson, fsion.ValueInt())
	if fsion.ValueInt() != -1 {
		t.Errorf("Test Get failed")
	}
}

func TestIsValidNumber(t *testing.T) {
	fsion := root.Get("")
	if !fsion.isValidNum("123") ||
		!fsion.isValidNum("-1.2") ||
		!fsion.isValidNum("1.2") ||
		fsion.isValidNum("1.34.223") ||
		fsion.isValidNum("+1.34223") ||
		fsion.isValidNum("abcad") {
		t.Errorf("Test Is valid Number failed")
	}
}

func TestValueBool(t *testing.T) {
	fsion := root.Get("bool")
	if !fsion.ValueBool() {
		t.Errorf("Test valueBool() failed")
	}
	if root.Get("bool1").ValueBool() {
		t.Errorf("Test valueBool() failed")
	}
	if !root.Get("bool2").ValueBool() {
		t.Errorf("Test valueBool() failed")
	}
	if root.Get("bool3").ValueBool() {
		t.Errorf("Test valueBool() failed")
	}
	if root.Get("k1").ValueBool() {
		t.Errorf("Test valueBool() failed")
	}
	if root.Get("str1").ValueBool() {
		t.Errorf("Test valueBool() failed")
	}
}
