/*
* @Author: 刘东旭
* @Date: 2020-11-22 12:23:14
 * @LastEditors: 刘东旭
 * @LastEditTime: 2020-11-26 13:03:10
* @Description:
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	//把key按abcd...排序，然后输出
	map1 := map[string]string{
		"username": "张三",
		"age":      "12",
		"sex":      "男",
		"height":   "150",
	}
	// 1.把key提取出来
	var mapKeys []string
	for key, _ := range map1 {
		mapKeys = append(mapKeys, key)
	}
	// 2.对key排序 abcdefg
	sort.Strings(mapKeys)
	for key, _ := range mapKeys {
		fmt.Printf("key=%v value=%v\n", mapKeys[key], map1[mapKeys[key]])
	}

}
