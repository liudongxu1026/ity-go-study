package main

import "fmt"

type Person struct {
	string
	int
}

func main() {
	p1 := Person{
		"张三",
		15,
	}
	// main.Person{string:"张三", int:15}
	fmt.Printf("%#v", p1)
}
