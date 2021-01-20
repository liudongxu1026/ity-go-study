/*
* @Author: 刘东旭
* @Date: 2020-11-12 12:47:05
 * @LastEditors: 刘东旭
 * @LastEditTime: 2020-11-12 13:17:00
* @Description:
*/
package main

import (
	"fmt"
	"strings"
)

func main() {

	// str1 := "第一个字符串"
	// str2 := "第二个字符串"
	// str3 := fmt.Sprintf("%v + %v", str1, str2)
	// fmt.Printf("%v", str3)
	// str1 := "a-b-c-d-e"
	// strSplit := strings.Split(str1, "-") //返回一个切片
	// fmt.Printf("%v", strSplit)           // [a b c d e]
	// strJoin := strings.Join(strSplit, "++") //返回一个字符串
	// fmt.Printf("%v", strJoin)               //a++b++c++d++e

	str1 := "啊哈哈啊哈"
	str2 := "啊"
	// fmt.Printf("%b", strings.HasPrefix(str1, str2)) //true str1是否以str2内容开头
	// fmt.Printf("%b", strings.HasSuffix(str1, str2)) //false str1是否以str2内容结尾
	fmt.Printf("%v\n", strings.Index(str1, str2))   // str2在str1出现的索引位置
	fmt.Printf("%v", strings.LastIndex(str1, str2)) // str2在str1结尾出现的索引位置(从左向右查找)
}
