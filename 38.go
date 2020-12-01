package main

import "fmt"

// 首字母大写 Person 表示公有结构体
// 首字母小写 person 表示私有结构体
type Person struct {
	name   string
	age    int
	sex    string
	height int
}

// 首字母大写 PrintInfo 表示公有方法
// 首字母小写 printInfo 表示私有方法
// func (谁 . 的这个方法，就是谁) 函数(参数){}
func (p Person) PrintInfo() {
	fmt.Printf("%#v", p)
}

// 可以接收指针类型参数
func (p *Person) SetName(name string) {
	p.name = name //通过指针修改 name
}
func main() {
	var p1 = Person{
		name: "赵四",
		age:  11,
		sex:  "女",
	}
	p1.SetName("王五")
	p1.PrintInfo() // main.Person{name:"王五", age:11, sex:"女", height:0}

}
