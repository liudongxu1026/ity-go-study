/*
* @Author: 刘东旭
* @Date: 2020-11-22 12:23:14
 * @LastEditors: 刘东旭
 * @LastEditTime: 2020-11-30 18:30:58
* @Description:
*/
package main

import "fmt"

func main() {
	// 使用new函数分配一个地址 int默认值是0
	// a存的new分配的地址 *a指针变量对应的值 &a取出a自己的地址并非new分配的地址
	a := new(int)
	fmt.Println(a, *a, &a)
}
