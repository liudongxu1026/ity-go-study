/*
* @Author: 刘东旭
* @Date: 2020-11-22 12:23:14
 * @LastEditors: 刘东旭
 * @LastEditTime: 2020-11-29 12:21:44
* @Description:
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	// 时间间隔类型常量
	fmt.Println(time.Nanosecond)  //1ns 1纳秒
	fmt.Println(time.Microsecond) //1µs 1微秒 1000*1纳秒
	fmt.Println(time.Millisecond) //1ms 1毫秒 1000*1微秒
	fmt.Println(time.Second)      //1s 1秒 1000*1毫秒
	fmt.Println(time.Minute)      //1m0s 1分钟
	fmt.Println(time.Hour)        //1h0m0s 1小时
}
