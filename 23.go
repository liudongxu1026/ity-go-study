/*
* @Author: 刘东旭
* @Date: 2020-11-22 12:23:14
 * @LastEditors: 刘东旭
 * @LastEditTime: 2020-11-29 11:45:59
* @Description:
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	nowTime := time.Now()
	// 2006表示年 01表示月 02表示日
	// (03表示时 12小时制) (15表示时 24小时制) 04分 05秒
	formatTime := nowTime.Format("03:04:05 2006-01-02")
	fmt.Printf("%v\n", nowTime)    //2020-11-29 11:39:55.4455227 +0800 CST m=+0.012962001
	fmt.Printf("%v\n", formatTime) // 11:39:55 2020-11-29
}
