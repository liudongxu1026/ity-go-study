/*
* @Author: 刘东旭
* @Date: 2020-11-22 12:23:14
 * @LastEditors: 刘东旭
 * @LastEditTime: 2020-11-29 11:51:53
* @Description:
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	nowTime := time.Now()
	unixTime := nowTime.Unix()
	unixNanoTime := nowTime.UnixNano()
	fmt.Println("时间戳", unixTime)
	fmt.Println("纳秒时间戳", unixNanoTime)
}
