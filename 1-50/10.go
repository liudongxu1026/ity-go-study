/*
* @Author: 刘东旭
* @Date: 2020-11-22 12:23:14
 * @LastEditors: 刘东旭
 * @LastEditTime: 2020-11-24 12:31:16
* @Description:
*/
package main

import "fmt"

func main() {
	var userInfo = make([]map[string]string, 3, 3)
	userInfo[0] = make(map[string]string)
	userInfo[0]["username"] = "张三"
	fmt.Printf("%v\n", userInfo[0])
	fmt.Printf("%v", userInfo)
}
