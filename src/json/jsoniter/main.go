package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

func main() {
	var data interface{}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var input = `
{
  "name": "I'm JSON",
  "string": "This is a string",
  "int": 1111111111111111111,
  "float": 3.141592653589757,
  "bool": false,
  "list": [1, 2, 3, "4"],
  "null": null,
  "dict": {
    "name": "I'm dick",
    "list": [
      "Stringer tinker bell",
      2.718281828459045,
      true,
      false,
      null,
      { "list": ["Funny little bulks", 2333], "blanks": "" }
    ],
    "object": { "AMD YES": "YES YES!!!", "INTEL": "SUCKS!!!" },
    "中文": "博大精深！"
  }
}
`
	json.Unmarshal([]byte(input), &data)
	fmt.Println(data)
}
