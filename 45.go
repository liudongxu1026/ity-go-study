package main

import "fmt"

type Animal struct {
	name string
}

func (a Animal) say() {
	fmt.Printf("%v是动物\n", a.name)
}

type Dog struct {
	age    int
	Animal // 结构体嵌套 继承
}

func (d Dog) wang() {
	fmt.Printf("%v在旺旺\n", d.name)
}

func main() {
	d1 := Dog{
		age: 8,
		Animal: Animal{
			name: "阿琪",
		},
	}
	d1.say()              // 阿琪是动物
	d1.wang()             // 阿琪在旺旺
	fmt.Printf("%#v", d1) // main.Dog{age:8, Animal:main.Animal{name:"阿琪"}}
}
