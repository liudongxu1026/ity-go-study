/*
* @Author: 刘东旭
* @Date: 2020-11-22 12:23:14
 * @LastEditors: 刘东旭
 * @LastEditTime: 2020-11-29 12:13:18
* @Description:
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	var timeStr string = "2020-11-29 11:51:57"
	var temp string = "2006-01-02 03:04:05"
	timeObj, _ := time.ParseInLocation(temp, timeStr, time.Local) //转成时间对象
	var unixTime int64 = timeObj.Unix()                           //转成时间戳
	fmt.Println(timeObj)                                          //2020-11-29 11:51:57 +0800 CST
	fmt.Println(unixTime)                                         //1606621917

}
