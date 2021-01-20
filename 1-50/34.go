/*
* @Author: 刘东旭
* @Date: 2020-11-22 12:23:14
 * @LastEditors: 刘东旭
 * @LastEditTime: 2020-11-30 18:21:01
* @Description:
*/
package main

import "fmt"

func main() {
	var userInfo = make(map[string]string)
	userInfo["name"] = "张三"
	fmt.Println(userInfo)

	var slice1 = make([]int, 4, 4)
	slice1[0] = 1
	fmt.Println(slice1) // [1 0 0 0]
}
