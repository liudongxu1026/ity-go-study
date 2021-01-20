/*
 * @Author: 刘东旭
 * @Date: 2020-11-22 12:23:14
 * @LastEditors: 刘东旭
 * @LastEditTime: 2020-11-23 12:58:17
 * @Description:  选择排序
 */
package main

import "fmt"

func main() {
	arr1 := []int{5, 2, 6, 99, 3, 8}
	for i := 0; i < len(arr1); i++ {
		for j := i + 1; j < len(arr1); j++ {
			fmt.Printf("%d %d\n", arr1[i], arr1[j])
			if arr1[i] > arr1[j] {
				num := arr1[i]
				arr1[i] = arr1[j]
				arr1[j] = num
			}
		}
	}
	// fmt.Printf("%v", arr1)
}
