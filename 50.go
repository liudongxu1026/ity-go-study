package main

import (
	"encoding/json"
	"fmt"
)

type Class struct {
	Title    string
	Students []Student
}
type Student struct {
	Id     int
	Name   string
	Gender string
}

func main() {
	class1 := Class{
		Title:    "数学",
		Students: make([]Student, 0, 200),
	}
	nameArr := []string{"张三", "李四", "王五", "赵六"}
	sex := []string{"男", "女"}
	for i, val := range nameArr {
		oneStu := Student{
			Id:     i,
			Name:   val,
			Gender: sex[i%2],
		}
		class1.Students = append(class1.Students, oneStu)
	}
	// 转json
	byteArr, _ := json.Marshal(class1)
	jsonStr := string(byteArr)
	fmt.Printf("%+v\n", class1)
	fmt.Printf("%s", jsonStr)
}
