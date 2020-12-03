package main

import (
	"encoding/json"
	"fmt"
)

type Animal struct {
	Name     string
	Age      int
	TypeName string
	EatWhat  string
	// sex      string //私有属性不能被json包访问
}

func main() {
	a1 := Animal{
		Name:     "大象",
		Age:      1,
		TypeName: "哺乳",
		EatWhat:  "香蕉",
	}
	// fmt.Printf("%#v", a1)
	// 1.把解构体转成 []uint8的切片
	jsonByte, _ := json.Marshal(a1)
	// 2.再用 string 就转成 json
	jsonStr := string(jsonByte)
	// {"Name":"大象","Age":1,"TypeName":"哺乳","EatWhat":"香蕉"}
	fmt.Printf("%v %T", jsonStr, jsonStr)
	// fmt.Printf("%v %T", jsonByte, jsonByte)
}
