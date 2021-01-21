package calc

import "fmt"

func init() {
	fmt.Println("add init")
}

// 公有方法
func Add(x, y int) int {
	return x + y
}
