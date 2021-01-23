package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}

// 电脑结构体
type Computer struct {
}

// Computer结构体的方法，接收参数表示 我要满足Usb接口形式的参数
func (c Computer) work(aaaa Usb) {
	aaaa.Start()
	aaaa.Stop()
}


// 手机结构体
type Phone struct{
	Name string
}
func (p Phone) Start(){
	fmt.Printf("Start %v\n", p.Name)
}
func (p Phone) Stop(){
	fmt.Printf("Stop %v\n", p.Name)
}

// 照相机结构体
type Camera struct{
	Name string
}
func (p Camera) Start(){
	fmt.Printf("Start %v\n", p.Name)
}
func (p Camera) Stop(){
	fmt.Printf("Stop %v\n", p.Name)
}


func main() {
	var computer1 Computer = Computer{}
	var phone1 Phone = Phone{
		Name: "华为手机",
	}
	var camera1 Camera = Camera{
		Name: "索尼相机",
	}
	computer1.work(phone1) // Start 华为手机 Stop 华为手机
	computer1.work(camera1) // Start 索尼相机 Stop 索尼相机
}
