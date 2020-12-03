/*
* @Author: 刘东旭
* @Date: 2020-11-12 12:47:05
 * @LastEditors: 刘东旭
 * @LastEditTime: 2020-12-03 13:07:24
* @Description:
*/
package main

func main() {

	// var a = 'a'
	// 1.直接输出 byte(字符) 的时候，输出的其实是这个字符对应码值
	// 值：97, 类型：int32
	// fmt.Printf("值：%v, 类型：%T", a, a)

	// 2.使用 %c 可以原样输出字符的值
	// 值：a, 类型：int32
	// fmt.Printf("值：%c, 类型：%T", a, a)
	// var i int = 20
	// str1 := fmt.Sprintf("%d", i)
	// fmt.Printf("%v, %T", str1, str1)

	// FormatInt(int64类型的值, int类型的进制)
	// str1 := strconv.FormatInt(int64(i), 10)
	// fmt.Printf("值：%v 类型：%T\n", str1, str1)

	// var f float32 = 27.67654765
	// str1 := strconv.FormatFloat(float64(f), 'f', 2, 64)
	// 值：27.68 类型：string
	// fmt.Printf("值：%v 类型：%T\n", str1, str1)

	// str := "1234"
	// fmt.Printf("%v--%T", str, str)
	// int1, _ := strconv.ParseInt(str, 10, 64)
	// fmt.Printf("%v--%T", int1, int1)
	// float1, _ := strconv.ParseFloat(str, 64)
	// fmt.Printf("%v--%T", float1, float1) //1234 float64

}
