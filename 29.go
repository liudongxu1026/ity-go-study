/*
* @Author: 刘东旭
* @Date: 2020-11-22 12:23:14
 * @LastEditors: 刘东旭
 * @LastEditTime: 2020-11-30 13:03:02
* @Description:
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	n := 5
	for t := range ticker.C {
		if n == 0 {
			ticker.Stop()
			break
		}
		n--
		fmt.Printf("%v\n", t)
	}
}
