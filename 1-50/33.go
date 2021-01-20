/*
* @Author: 刘东旭
* @Date: 2020-11-22 12:23:14
 * @LastEditors: 刘东旭
 * @LastEditTime: 2020-11-30 18:12:32
* @Description:
*/
package main

import "fmt"

func fn1(a int) {
	a = 10
}

func fn2(a *int) {
	*a = 20
}

func main() {
	var a = 5
	fn1(a)
	fmt.Println(a) //5
	fn2(&a)
	fmt.Println(a) //20
}
