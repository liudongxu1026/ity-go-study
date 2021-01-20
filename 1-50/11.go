/*
* @Author: 刘东旭
* @Date: 2020-11-22 12:23:14
 * @LastEditors: 刘东旭
 * @LastEditTime: 2020-11-25 11:42:42
* @Description:
*/
package main

import "fmt"

func main() {
	var map1 = make(map[int]int, 10)
	map1[3] = 10
	map1[2] = 8
	map1[8] = 14
	map1[81] = 14
	map1[9] = 5
	map1[4] = 5
	fmt.Printf("%v", map1)
}
