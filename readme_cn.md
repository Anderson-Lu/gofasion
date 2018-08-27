GoFasion:一个轻量级的具备链式调用风格的JSON数据解析神器
---

[English Document](https://github.com/Anderson-Lu/gofasion/blob/master/readme.md)

Gofasion是一个方便开发过程中接口JSON数据解析的轻量级解析库，其最大的特点在于支持链式调用，也就是说不必预先定义好数据的结构就可以直接获取到目标键名和键值。

### 开源

[https://github.com/Anderson-Lu/gofasion](https://github.com/Anderson-Lu/gofasion)

### 安装

```shell
$ go get github.com/Anderson-Lu/gofasion/gofasion
```

### Go module

```shell
//go.mod
module github.com/Anderson-Lu/gofasion/gofasion
```

### 如何定位JSON节点

你可以把一个JSON数据想象成一颗树，每个元素都是树上的节点(*Fastion),节点的值可以是任意类型(bool,int etc.)，通过链式调用可以从根节点追溯到任意节点，从而取出其中的数值。

```json
{
  "level1":{
      "level2":{
          "level3":1
        }
    }
}
```

要取出上面的`level3`的值，可以通过下面的例子快速访问：

```golang
fsion := gofasion.NewFasion(yourJsonStr)
level3 := fsion.Get("level1").Get("level2").Get("level3").ValueStr()

//或者直接使用GetFromPath(dir string)方法直接访问
//level3 := fsion.GetFromPath("level1.level2.level3").ValueStr()
```

### 如何遍历JSON数组

我们提供了`Array()`方法，用于表示数组类型的JSON数据，对于数组中的元素，则又是一个`*Fasion`对象，仍然可以进行上述操作。

```json
{
  "array" : [
    {"name":1},
    {"name":2}
  ]
}
```

要遍历`array`中的数据，只需要这样做：

```golang
array := fsion.Get("array").Array()
for _,v := range array{
  name := v.Get("name").ValueInt()
  //your code
}
```

### 如何遍历不规则JSON数据或者无键名数据

很多时候，我们需要解析不规则的JSON数据，比如:

```json
[
  1,2,"helloword",{"name":"demo"}
] 
```

可以通过`Array()`方法，快速取值

### 快速开始

```golang
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
  
  //输出 {"name":"foo","value":1...}
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

### 基本方法

```golang

  //how to create *Fasion instance
  func NewFasion(rawJson string) *Fasion               //Create Fasion From raw json
  func NewFasionFromBytes(rawJson []byte) *Fasion      //Create Fasion From bytes
  func NewFasionFromUrl(targetUrl string, params url.Values) *Fasion  //Create Fasion From http get

  //Methods for *Fasion
  Get(key string) *IFasion          //获取JSON节点对象,每个节点对象包含下面所有方法
  GetFromPath(dir string) *IFasion  //获取JSON节点对象(路径比如a.b.c)

  //Methods to get value from *Fasion node
  ValueStr() string         //获取节点的字符串值
  ValueInt() int            //获取节点的int值
  ValueInt16() int16 
  ValueInt32() int32
  ValueInt64() int64
  ValueFloat32() float32
  ValueFloat64() float64
  ValueBool() bool
  Array() []*Fasion         //获取节点的数组对象
  ArrayForEach(func(int, *Fasion)) //直接遍历对象数组
  Value(interface{}) error  //与json.Marshal()类似
```

### 性能

1,000,000 `*Fastion.Get()` 耗费 5,000ms ~ 7,000ms.

### 版本

`v1` 基础版本，提供常用的基本功能

### 贡献

欢迎大家提出宝贵issue，也可以提交合并请求，希望能做一款让所有golang开发者收益的开源库。

### 许可

MIT Licence