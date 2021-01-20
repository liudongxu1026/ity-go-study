/*
* @Author: 刘东旭
* @Date: 2020-11-22 12:23:14
 * @LastEditors: 刘东旭
 * @LastEditTime: 2020-12-03 18:45:37
* @Description:
*/
package main

import "fmt"

// 自定义一个方法类型
type calcType func(int, int) int

func calc(x, y int, cFn calcType) int {
	return cFn(x, y)
}

func add(x, y int) int {
	return x + y
}

func main() {
	sum := calc(10, 5, add)
	fmt.Printf("%v", sum)
}
