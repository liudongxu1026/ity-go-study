/*
* @Author: 刘东旭
* @Date: 2020-11-22 12:23:14
 * @LastEditors: 刘东旭
 * @LastEditTime: 2020-11-27 13:56:34
* @Description:
*/
package main

import "fmt"

func adder1() func() int {
	x := 10
	return func() int {
		return x + 1
	}
}
func adder2() func(int) int {
	x := 10
	return func(y int) int {
		x += y
		return x
	}
}
func main() {
	var fn1 = adder1()        //此方法返回一个函数
	fmt.Printf("%v\n", fn1()) //11
	fmt.Printf("%v\n", fn1()) //11
	fmt.Printf("%v\n", fn1()) //11

	var fn2 = adder2()          //此方法返回一个函数
	fmt.Printf("%v\n", fn2(10)) //20
	fmt.Printf("%v\n", fn2(20)) //40
	fmt.Printf("%v\n", fn2(30)) //70

}
