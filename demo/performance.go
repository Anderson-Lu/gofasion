package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Anderson-Lu/gofasion/gofasion"
)

//this file is written to test the performance of gofasion
//gofssion is useful to get specific keys but if you want to unmarshal all struct, please use json.Unmarshal instead.
func Test1() {
	testjson := "{\"key1\":\"value1\",\"key2\":\"value2\"}"
	fasion := gofasion.NewFasion(testjson)
	begin := time.Now().UnixNano() / 1000000
	var s string
	for i := 0; i < 1000000; i++ {
		s = fasion.Get("key1").ValueStr()
	}
	fmt.Println(s)
	fmt.Println("1,000,000 GET Operation Costs ", time.Now().UnixNano()/1000000-begin, "ms")
	var tmp struct {
		Key1 string `json:"key1"`
	}
	begin = time.Now().UnixNano() / 1000000
	for i := 0; i < 1000000; i++ {
		json.Unmarshal([]byte(testjson), &tmp)
		s = tmp.Key1
	}
	fmt.Println("1,000,000 Json.Unmarshal Operation Costs ", time.Now().UnixNano()/1000000-begin, "ms")
}
