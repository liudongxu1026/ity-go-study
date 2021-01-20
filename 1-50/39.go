package main

import "fmt"

// 注：非本地类型不能定义方法，不能给别的包的类型定义方法
type MyInt int

func (m MyInt) printInfo() {
	fmt.Println("自定义类型的自定义方法")
}

func main() {
	var a MyInt = 20
	a.printInfo()
}
