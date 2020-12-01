package main

import "fmt"

// 首字母大写 Person 表示公有结构体
// 首字母小写 person 表示私有结构体
type Person struct {
	name string
	age  int
	sex  string
}

func main() {
	var p1 = Person{
		name: "赵四",
		age:  11,
		sex:  "女",
	}
	p2 := p1
	p2.age = 999
	fmt.Printf("%#v\n", p1)
	fmt.Printf("%#v", p2)

}
