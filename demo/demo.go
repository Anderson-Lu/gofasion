package main

import (
	"fmt"

	"github.com/Anderson-Lu/gofasion/gofasion"
)

func DemoGetKeys() {
	fsion := gofasion.NewFasion(`{"name":"hello","value":"1111"}`)
	fmt.Println(fsion.Keys())
}

func DemoHasKeys() {
	fsion := gofasion.NewFasion(`{"name":"hello","value":"1111"}`)
	fmt.Println(fsion.HasKey("name"))
	fmt.Println(fsion.HasKey("value1"))
}

func DemoCheckJsonFormat() {
	fsion := gofasion.NewFasion("i am a wrong json encoding format")
	fmt.Println(fsion.IsValidJson())
}

func DemoParseStringVal() {
	fsion := gofasion.NewFasion(`{"name":"hello world & hello world"}`)
	fmt.Println(fsion.Get("name").ValueStr())
}

func DemoParseIntVal() {
	fsion := gofasion.NewFasion(`{"name":1}`)
	fmt.Println(fsion.Get("name").ValueInt())
}

func DemoParseFloat64() {
	fsion := gofasion.NewFasion(`{"name":1.000001}`)
	fmt.Println(fsion.Get("name").ValueFloat64())
}

func DemoParseArray() {
	fsion := gofasion.NewFasion(`[1,2,3,45]`)
	fmt.Println(fsion.Array())
}

func DemoParseArrayForEach() {
	fsion := gofasion.NewFasion(`[1,2,3,4,5]`)
	fsion.ArrayForEach(func(idx int, node *gofasion.Fasion) {
		fmt.Println(idx, node.ValueInt())
	})
}

var testJson = `
	{
		"name":"foo",
		"value":1.1,
		"second_level": {"name":2},
		"second_array":[1,2,3,4,5,6,7],
		"bool": true,
		"value64":1234567890
	}
`

var testJson2 = `
  [
	  1.0,2,"helloword",{"name":"demo"}
  ]  
`

func main() {

	fsion := gofasion.NewFasion(testJson)
	fmt.Println(fsion.Get("name").ValueStr())
	fmt.Println(fsion.Get("value").ValueInt())
	fmt.Println(fsion.Get("value").ValueFloat64())
	fmt.Println(fsion.Json())

	return

	i32 := fsion.Get("value").ValueInt32()
	fmt.Println(i32)

	i64 := fsion.Get("value64").ValueInt64()
	fmt.Println(i64)

	second_fson := fsion.Get("second_level")
	fmt.Println(second_fson.Get("name").ValueStr())

	second_array := fsion.Get("second_array").Array()
	for _, v := range second_array {
		fmt.Println(v.ValueInt())
	}

	fmt.Println(fsion.GetFromPath("second_level.name").ValueStr())

	boolVal := fsion.Get("bool").ValueStr()
	fmt.Println(boolVal)

	fsion2 := gofasion.NewFasion(testJson2)
	elems := fsion2.Array()
	fmt.Println(elems[0].ValueInt())
	fmt.Println(elems[1].ValueInt())
	fmt.Println(elems[2].ValueStr())

	fmt.Println(elems[3].Json())

	var iter struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	}
	fsion.Value(&iter)
	fmt.Println(iter.Name)
	fmt.Println(iter.Value)

	//performance test
	// Test1()
}

func demo2() {

	yourJson := `
		"data": {
			"user": {
				"user_id": "",
				"email": "",
				"name": "",
				"surname": "",
				"telephone": "",
				"country": "",
				"time_zone": "",
				"date_format": "DD/MM/YYYY 24H",
				"language": "en",
				"company_default": "",
				"status": "not confirmed",
				"companies": [
					{
					"data1":""
					"data2":""
					}
				],
				"superadmin": "",
				"accept_newsletter": false
			}
		},
		"method": "POST",
		"url": "/v1.0/url/test",
		"date": "2018-08-09T11:00:55+02:00",
		"status": "SUCCESS",
		"code": "200",
		"message": "HTTP OK",
		"details": "Valid"
	}
	`

	fsion := gofasion.NewFasion(yourJson)
	data := fsion.Get("data")
	user := data.Get("user")

	userId := user.Get("user_id").ValueStr()
	email := user.Get("email").ValueStr()

	fmt.Println(userId, email)

	companies := user.Get("companies").Array()

	for _, v := range companies {
		fmt.Println(v.Get("data1").ValueStr())
		fmt.Println(v.Get("data2").ValueStr())
	}

}
