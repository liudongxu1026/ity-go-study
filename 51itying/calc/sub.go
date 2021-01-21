package calc

import "fmt"

func init() {
	fmt.Println("sub init")
}

// 公有方法
func Sub(x, y int) int {
	return x - y
}
