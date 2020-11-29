/*
* @Author: 刘东旭
* @Date: 2020-11-22 12:23:14
 * @LastEditors: 刘东旭
 * @LastEditTime: 2020-11-29 12:23:50
* @Description:
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	var timeObj = time.Now()
	// 2020-11-29 12:23:24.4083416 +0800 CST m=+0.014961101
	fmt.Printf("%v\n", timeObj)
	timeObj = timeObj.Add(time.Hour) //加一小时
	// 2020-11-29 13:23:24.4083416 +0800 CST m=+3600.014961101
	fmt.Printf("%v\n", timeObj)

}
