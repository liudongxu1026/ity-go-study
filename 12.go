/*
* @Author: 刘东旭
* @Date: 2020-11-22 12:23:14
 * @LastEditors: 刘东旭
 * @LastEditTime: 2020-11-25 11:56:10
* @Description:
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	var words = "how do you do how do you do how do you do"
	var splitWords = strings.Split(words, " ")
	var time = make(map[string]int)
	for _, value := range splitWords {
		_, isKeyInTime := time[value]
		if isKeyInTime {
			time[value]++
		} else {
			time[value] = 1
		}
	}
	fmt.Printf("%v", time)
}
