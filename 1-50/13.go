/*
* @Author: 刘东旭
* @Date: 2020-11-22 12:23:14
 * @LastEditors: 刘东旭
 * @LastEditTime: 2020-11-25 13:11:40
* @Description:
*/
package main

import "fmt"

// x,y int 表示 x和y都是int类型
// func fn(x, y int) int {
// 	return x - y
// }

// ...表示参数数量不固定，接收时会自动转为切片
// func fn2(x ...int) {
// 	fmt.Printf("%v", x)
// }

// 第一个参数会赋值给x, y会将后面所有参数转换为切片
// func fn3(x int, y ...int) {
// 	fmt.Printf("x=%v, y=%v", x, y)
// }

// 多个返回值
// func fn4(x, y int) (int, int) {
// 	return x + y, x * y
// }
// 多个返回值 返回值命名
func fn5(x, y int) (a int, b int) {
	a = x + y
	b = x * y
	return //使用return关键字直接返回
}

func main() {
	// fn2(1, 2, 3, 4, 5, 6)
	// fn3(100, 1, 3, 5, 7, 9)
	// num1, num2 := fn4(5, 3)
	// fmt.Printf("%d, %d", num1, num2)

	aa, bb := fn5(5, 3)
	fmt.Printf("%d, %d", aa, bb) //8, 15

}
