package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Address Address
}

type Address struct {
	Name  string
	Phone int
	City  string
}

func main() {
	p1 := Person{
		Name: "张三",
		Age:  10,
		Address: Address{
			Name:  "新城子",
			Phone: 15212341234,
			City:  "沈阳",
		},
	}
	fmt.Printf("%#v\n", p1)
}
