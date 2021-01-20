/*
 * @Author: 刘东旭
 * @Date: 2020-11-22 12:23:14
 * @LastEditors: 刘东旭
 * @LastEditTime: 2020-11-22 12:27:22
 * @Description:
 */
package main

import "fmt"

func main() {
	sliceA := []int{1, 2, 3, 4, 5}
	sliceB := make([]int, 5, 5)
	fmt.Printf("复制前：sliceA: %v\n", sliceA) //[1 2 3 4 5]
	fmt.Printf("复制前：sliceB: %v\n", sliceB) //[0 0 0 0 0]
	// 把sliceA复制到sliceB中
	copy(sliceB, sliceA)
	fmt.Printf("复制后：sliceA: %v\n", sliceA) //[1 2 3 4 5]
	fmt.Printf("复制后：sliceB: %v\n", sliceB) //[1 2 3 4 5]
	sliceA[0] = 11
	sliceB[0] = 111
	fmt.Printf("修改后：sliceA: %v\n", sliceA) //[11 2 3 4 5]
	fmt.Printf("修改后：sliceB: %v\n", sliceB) //[1 2 3 4 5]
}
