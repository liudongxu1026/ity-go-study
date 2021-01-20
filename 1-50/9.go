/*
* @Author: 刘东旭
* @Date: 2020-11-22 12:23:14
 * @LastEditors: 刘东旭
 * @LastEditTime: 2020-11-24 12:16:01
* @Description:  
*/
package main

import "fmt"

func main() {
	var userInfo = map[string]string{
		"name": "zhangsan",
		"age":  "19",
		"sex":  "女",
	}
	// userInfo["sex"] = "男"
	// v是获取到的值 ok表示 这个key是否存在
	// v, ok := userInfo["hhh"]
	// fmt.Printf("v=%v 是否存在=%v", v, ok)

	// for k, v := range userInfo {
	// 	fmt.Printf("key=%s value=%s\n", k, v)
	// }
	delete(userInfo, "sex")
	fmt.Printf("%+v", userInfo)
}
