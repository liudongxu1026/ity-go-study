/*
* @Author: 刘东旭
* @Date: 2020-11-22 12:23:14
 * @LastEditors: 刘东旭
 * @LastEditTime: 2020-12-04 14:16:07
* @Description:
*/
package main

import "fmt"

func main() {
	// 使用new函数分配一个地址 int默认值是0
	// a 存的new分配的地址
	// *a 指针变量对应的值
	// &a 取出a自己的地址并非new分配的地址
	a := new(int)
	// 0xc000012090 0 0xc000006028 0xc000012090
	// *a = "123" // 不可以 因为 new 分配的地址类型是 int
	*a = 1 // √
	fmt.Println(a, *a, &a, *&a)
}
