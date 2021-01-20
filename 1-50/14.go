/*
* @Author: 刘东旭
* @Date: 2020-11-22 12:23:14
 * @LastEditors: 刘东旭
 * @LastEditTime: 2020-11-25 19:03:12
* @Description:
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	var words = "what do you want to do what HOW DO YOU LIKE IT I LIKE"
	var wordSplit = strings.Split(words, " ")
	var num = make(map[string]int)
	for _, value := range wordSplit {
		theKey := strings.ToUpper(value)
		_, isInNum := num[theKey]
		if isInNum {
			num[theKey]++
		} else {
			num[theKey] = 1
		}
	}
	fmt.Printf("最终结果: %v", num)
}
