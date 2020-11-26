/*
 * @Author: 刘东旭
 * @Date: 2020-11-22 12:23:14
 * @LastEditors: 刘东旭
 * @LastEditTime: 2020-11-22 12:42:10
 * @Description:
 */
package main

import "fmt"

func main() {
	// 1.扩容
	// sliceA := []int{1, 2, 3, 4, 5}
	// sliceB := append(sliceA, 6)
	// fmt.Printf("%v 长度%d 容量%v", sliceB, len(sliceB), cap(sliceB))
	// 2.删除
	// 删除索引为 2 的数值
	sliceA := []int{1, 2, 3, 4, 5}
	// sliceA[:2] 从头取到第二位, sliceA[3:] 从第三位一直到结尾
	// ...在前表示合并为切片，...在后表示解散
	sliceA = append(sliceA[:2], sliceA[3:]...)
	fmt.Printf("%v", sliceA) //[1 2 4 5]
}
