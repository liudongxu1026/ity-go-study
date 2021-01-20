package main

import "fmt"

type Person struct {
	name  string
	age   int
	hobby []string //这是一个切片
	map1  map[string]string
}

func main() {
	p1 := Person{
		name: "张三",
		age:  10,
		// hobby: []string{"抽烟", "喝酒", "烫头"},
		hobby: make([]string, 3, 5), //分配存储空间，如果不分配的话，零值是 nil
		map1: map[string]string{
			"a": "1",
			"b": "2",
			"c": "3",
		},
		// map1:  make(map[string]string), //分配存储空间，如果不分配的话，零值是 nil
	}
	// main.Person{name:"张三", age:10, hobby:[]string{"", "", ""}, map1:map[string]string{"a":"1", "b":"2", "c":"3"}}
	fmt.Printf("%#v\n", p1)
}
