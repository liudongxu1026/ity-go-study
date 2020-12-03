package main

import "fmt"

type Person struct {
	name    string
	age     int
	address Address
}

type Address struct {
	name  string
	phone int
	city  string
}

func main() {
	p1 := Person{
		name: "张三",
		age:  10,
		address: Address{
			name:  "新城子",
			phone: 15212341234,
			city:  "沈阳",
		},
	}
	// main.Person{name:"张三", age:10, address:main.Address{name:"新城子", phone:15212341234, city:"上海"}}
	fmt.Printf("%#v\n", p1)
}
