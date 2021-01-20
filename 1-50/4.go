/*
* @Author: 刘东旭
* @Date: 2020-11-12 12:47:05
 * @LastEditors: 刘东旭
 * @LastEditTime: 2020-11-19 12:41:30
* @Description:
*/
package main

import "fmt"

func main() {
	// 普通循环
	// for i := 0; i < 5; i++ {
	// 	fmt.Printf("i=%v\n", i)
	// 	for j := 0; j < 3; j++ {
	// 		if j == 1 {
	// 			break
	// 		}
	// 		fmt.Printf("j=%v\n", j)
	// 	}
	// }
	// label循环
label:
	for i := 0; i < 5; i++ {
		fmt.Printf("i=%v\n", i)
		for j := 0; j < 3; j++ {
			if j == 1 {
				break label
			}
			fmt.Printf("j=%v\n", j)
		}
	}
}
