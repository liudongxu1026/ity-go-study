/*
* @Author: 刘东旭
* @Date: 2020-11-22 12:23:14
 * @LastEditors: 刘东旭
 * @LastEditTime: 2020-11-29 12:05:14
* @Description:
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	var unixTime int64 = 1606621917
	timeObj := time.Unix(unixTime, 0)
	var formatTime string = timeObj.Format("2006-01-02 03:04:05")
	fmt.Printf("%v\n", timeObj)    // 2020-11-29 11:51:57 +0800 CST
	fmt.Printf("%v\n", formatTime) // 2020-11-29 11:51:57
}
