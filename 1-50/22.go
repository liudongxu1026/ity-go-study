/*
* @Author: 刘东旭
* @Date: 2020-11-22 12:23:14
 * @LastEditors: 刘东旭
 * @LastEditTime: 2020-11-27 14:23:02
* @Description:
*/
package main

import "fmt"

func fn1() {
	fmt.Println("fn1")
}
func fn2() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err", err)
		}
	}()
	panic("抛出一个异常")
}

func main() {
	fn1()
	fn2()
	fmt.Println("结束")
}
