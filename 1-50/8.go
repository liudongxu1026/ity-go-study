/*
* @Author: 刘东旭
* @Date: 2020-11-22 12:23:14
 * @LastEditors: 刘东旭
 * @LastEditTime: 2020-11-26 12:04:43
* @Description:  冒泡排序
*/
package main

import "fmt"

func main() {
	arr1 := []int{5, 2, 6, 99, 3, 8}
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr1)-1-i; j++ {
			if arr1[j] > arr1[j+1] {
				// num := arr1[j]
				// arr1[j] = arr1[j+1]
				// arr1[j+1] = num
				arr1[j], arr1[j+1] = arr1[j+1], arr1[j]
			}
		}
	}
	fmt.Printf("%v", arr1)
}
