/*
* @Author: 刘东旭
* @Date: 2020-11-22 12:23:14
 * @LastEditors: 刘东旭
 * @LastEditTime: 2020-11-26 15:21:54
* @Description:
*/
package main

import "fmt"

// 定义一个calc类型
type calc func(int, int) int

func add(a int, b int) int {
	return a + b
}
func test1(a int, b int) int {
	return a + b
}
func main() {
	var fn1 calc = add
	fmt.Printf("%T\n", fn1)   // main.calc
	fmt.Printf("%T\n", test1) // func(int, int) int
}
