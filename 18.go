/*
* @Author: 刘东旭
* @Date: 2020-11-22 12:23:14
 * @LastEditors: 刘东旭
 * @LastEditTime: 2020-11-27 12:01:31
* @Description:
*/
package main

import "fmt"

// 递归实现
func addFn(num int) int {
	if num > 1 {
		return num + addFn(num-1)
	} else {
		return 1
	}
}

func main() {
	result := addFn(100)
	fmt.Printf("%v", result)

}
