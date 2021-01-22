package main

import "fmt"

// 接口
type Usb interface {
	start()
	stop()
}

// 如果接口里面有方法的话，必须要通过结构体或者通过自定义类型实现这个接口
type Phone struct {
	Name string
}

// Phone 结构体要实现 Usb 接口的话，必须得实现 Usb 接口中的所有方法
func (thisPhone Phone) start() {
	fmt.Printf("start %v\n", thisPhone.Name)
}
func (thisPhone Phone) stop() {
	fmt.Printf("stop %v\n", thisPhone.Name)
}

func main() {
	// golang中接口就是一个数据类
	// 表示实现Usb接口
	var p1 Usb = Phone{
		Name: "华为手机",
	}
	p1.start()
	p1.stop()
}
