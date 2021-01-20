package main

import (
	"51itying/calc"
	"fmt"
)

func main() {
	a1 := calc.Add(1, 10)
	a2 := calc.Sub(10, 2)
	fmt.Printf("%d\n", a1) // 11
	fmt.Printf("%d\n", a2) //8
}
