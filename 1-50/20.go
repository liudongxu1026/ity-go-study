/*
* @Author: 刘东旭
* @Date: 2020-11-22 12:23:14
 * @LastEditors: 刘东旭
 * @LastEditTime: 2020-12-18 16:33:57
* @Description:
*/
package main

import "fmt"

func fn1() int {
	var a int
	defer func() {
		a++
	}()
	return a
}
func fn2() (a int) {
	defer func() {
		a++
	}()
	return a
}
func main() {
	fmt.Println(fn1()) //0
	fmt.Println(fn2()) //1
}
