GoFasion: A lightweight JSON data parsing library with chained calling style
---

[中文文档](https://github.com/Anderson-Lu/gofasion/blob/master/readme_cn.md)

Gofasion is a lightweight parsing library that facilitates the parsing of interface JSON data during development. Its biggest feature is to support chained calls, which means that the target key name and key value can be directly obtained without pre-defining the structure of the data.


### Open source

[https://github.com/Anderson-Lu/gofasion](https://github.com/Anderson-Lu/gofasion)

### Required

go 1.9 or above is requried.

### Installation

```shell
$ go get github.com/Anderson-Lu/gofasion/gofasion
```

### Go module

```shell
//go.mod
module github.com/Anderson-Lu/gofasion/gofasion
```

### How to locate a JSON node

You can think of a JSON data as a tree, each element is a node on the tree (*Fastion), the value of the node can be any type (bool, int etc.), which can be traced from the root node through chained calls. Any node to take out the values.

```json
{
  "level1":{
      "level2":{
          "level3":1
        }
    }
}
```

To retrieve the value of `level3` above, you can quickly access it by the following example:

```golang
fsion := gofasion.NewFasion(yourJsonStr)
level3 := fsion.Get("level1").Get("level2").Get("level3").ValueStr()

//or fetch specific value by GetFromPath(dir string) method 
//level3 := fsion.GetFromPath("level1.level2.level3").ValueStr()
```

### How to traverse JSON arrays

We provide the `Array()` method to represent the JSON data of the array type. For the elements in the array, it is a `*Fasion` object, which can still be used.

```json
{
  "array" : [
    {"name":1},
    {"name":2}
  ]
}
```

To traverse the data in `array`, just do this:

```golang
array := fsion.Get("array").Array()
for _,v := range array{
  name := v.Get("name").ValueInt()
  //your code
}
```

### How to traverse irregular JSON data or no key name data

Many times, we need to parse irregular JSON data, such as:

```json
[
  1,2,"helloword",{"name":"demo"}
] 
```

Can quickly get values ​​by `Array()` method

### Quick start

```golang
package main

import (
  "github.com/Anderson-Lu/gofasion/gofasion"
  "fmt"
)

//Rule data
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

//Irregular data
var testJson2 = `
  [
    1,2,"helloword",{"name":"demo"}
  ]  
`

func main() {
  
  fsion := gofasion.NewFasion(testJson)

  //output "foo"
  fmt.Println(fsion.Get("name").ValueStr())
  
  //output 1
  fmt.Println(fsion.Get("value").ValueInt())
  
  //output {"name":"foo","value":1...}
  fmt.Println(fsion.Json())

  i32 := fsion.Get("value").ValueInt32()
  fmt.Println(i32)

  i64 := fsion.Get("value64").ValueInt64()
  fmt.Println(i64)

  second_fson := fsion.Get("second_level")
  fmt.Println(second_fson.Get("name").ValueStr())

  //Traversal of array data
  second_array := fsion.Get("second_array").Array()
  for _, v := range second_array {
    fmt.Println(v.ValueInt())
  }

  boolVal := fsion.Get("bool").ValueStr()
  fmt.Println(boolVal)

  //Analysis of irregular data
  fsion2 := gofasion.NewFasion(testJson2)
  elems := fsion2.Array()
  fmt.Println(elems[0].ValueInt())
  fmt.Println(elems[1].ValueInt())
  fmt.Println(elems[2].ValueStr())

  fmt.Println(elems[3].Json())

  //Traditional structure analysis
  var iter struct {
    Name  string `json:"name"`
    Value int    `json:"value"`
  }
  fsion.Value(&iter)
  fmt.Println(iter.Name)
  fmt.Println(iter.Value)
}

```

### Performance

1,000,000 `*Fastion.Get()` cost about 5,000ms ~ 7,000ms.

### Basic methods

```golang
  //how to create *Fasion instance
  func NewFasion(rawJson string) *Fasion               //Create Fasion From raw json
  func NewFasionFromBytes(rawJson []byte) *Fasion      //Create Fasion From bytes
  func NewFasionFromUrl(targetUrl string, params url.Values) *Fasion  //Create Fasion From http get

  //Methods for *Fasion
  Get(key string) *IFasion  //Get the JSON node object, each node object contains all the methods below
  GetFromPath(dir string) *IFasion //Get the JSON node via node path like node1.node2.node3

  //Methods to get value from *Fasion node
  ValueStr() string         //Get the string value of the node
  ValueInt() int            //Get the int value of the node
  ValueInt16() int16 
  ValueInt32() int32   
  ValueInt64() int64
  ValueFloat32() float32
  ValueFloat64() float64
  ValueBool() bool
  Array() []*Fasion         //Get the array object of the node
  ArrayForEach(func(int, *Fasion)) //Get the array object of the node iterator
  Value(interface{}) error  //Similar to json.Marshal()
  Json() string             //Get the JSON string of the node
  Keys() []string           //Get all keys of the node
  HasKey(string) bool       //Judge if the specific node contains specific key
```

### Version

`v1` Basic version, providing common basic functions

### Contribution

You are welcome to submit a valuable issue, you can also submit a merger request, hoping to make an open source library for all golang developers.

### License

MIT License
