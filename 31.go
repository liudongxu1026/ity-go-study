/*
* @Author: 刘东旭
* @Date: 2020-11-22 12:23:14
 * @LastEditors: 刘东旭
 * @LastEditTime: 2020-11-30 13:17:05
* @Description:
*/
package main

import "fmt"

func main() {
	var a = 10
	var p = &a //指针变量 p的类型=*int
	fmt.Printf("a=%v 类型=%T 地址=%p\n", a, a, &a)
	fmt.Printf("p=%v 类型=%T 地址=%p", p, p, &p)
}
