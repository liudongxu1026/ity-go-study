package main

import (
	"encoding/json"
	"fmt"
)

type Animal struct {
	Name     string `json:"id"`
	Age      int    `json:"age"`
	TypeName string `json:"typeName"`
	EatWhat  string `json:"eatWhat"`
}

func main() {
	a1 := Animal{
		Name:     "长颈鹿",
		Age:      2,
		TypeName: "哺乳",
		EatWhat:  "草料",
	}
	byte1, _ := json.Marshal(a1)
	json1 := string(byte1)
	// {"id":"长颈鹿","age":2,"typeName":"哺乳","eatWhat":"草料"}
	fmt.Println(json1)
}
