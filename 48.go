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
}

func main() {
	str1 := `{"Name":"大象","Age":1,"TypeName":"哺乳","EatWhat":"香蕉"}`
	s1 := Animal{}
	// 1.json.Unmarshal(字符切片, 结构体实例化地址(转换后会赋值给这个地址的变量))
	err := json.Unmarshal([]byte(str1), &s1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v", s1)
}
