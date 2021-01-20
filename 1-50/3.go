/*
* @Author: 刘东旭
* @Date: 2020-11-12 12:47:05
 * @LastEditors: 刘东旭
 * @LastEditTime: 2020-11-17 13:02:51
* @Description:
*/
package main

import "fmt"

func main() {
	for i := 0; i < 9; i++ {
		for i2 := 0; i2 < (i + 1); i2++ {
			fmt.Printf("%v*%v=%v ", i2+1, i+1, (i+1)*(i2+1))
		}
		fmt.Printf("\n")
	}

}
