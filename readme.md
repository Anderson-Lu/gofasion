GoFasion:一个轻量级的具备链式调用风格的JSON数据解析神器
---

Gofasion是一个方便开发过程中接口JSON数据解析的轻量级解析库，其最大的特点在于支持链式调用，也就是说不必预先定义好数据的结构就可以直接获取到目标键名和键值。

### 安装

```shell
go get github.com/Anderson-Lu/fasion/gofasion
```

### 快速开始

```shell
package main

import (
	"github.com/Anderson-Lu/fasion/gofasion"
	"fmt"
)

//规则数据
var testJson = `
	{
		"name":"foo",
		"value":1,
		"second_level": {"name":2},
		"second_array":[1,2,3,4,5,6,7],
		"bool": true,
		"value64":1234567890
	}
`

//不规则数据
var testJson2 = `
  [
	  1,2,"helloword",{"name":"demo"}
  ]  
`

func main() {
	fsion := gofasion.NewFasion(testJson)

    //输出 "foo"
	fmt.Println(fsion.Get("name").ValueStr())
	
    //输出 1
    fmt.Println(fsion.Get("value").ValueInt())
	
    //输出 {\"name\":\"foo\",\"value\":1...}
    fmt.Println(fsion.Json())

	i32 := fsion.Get("value").ValueInt32()
	fmt.Println(i32)

	i64 := fsion.Get("value64").ValueInt64()
	fmt.Println(i64)

	second_fson := fsion.Get("second_level")
	fmt.Println(second_fson.Get("name").ValueStr())

    // 数组数据的遍历
	second_array := fsion.Get("second_array").Array()
	for _, v := range second_array {
		fmt.Println(v.ValueInt())
	}

	boolVal := fsion.Get("bool").ValueStr()
	fmt.Println(boolVal)

    //不规则数据的解析
	fsion2 := gofasion.NewFasion(testJson2)
	elems := fsion2.Array()
	fmt.Println(elems[0].ValueInt())
	fmt.Println(elems[1].ValueInt())
	fmt.Println(elems[2].ValueStr())

	fmt.Println(elems[3].Json())

    //传统结构体解析
	var iter struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	}
	fsion.Value(&iter)
	fmt.Println(iter.Name)
	fmt.Println(iter.Value)
}

```

### 版本

`v1` 基础版本，提供常用的基本功能

### 贡献

欢迎大家提出宝贵issue，也可以提交合并请求，希望能做一款让所有golang开发者收益的开源库。

### 许可

MIT Licence
