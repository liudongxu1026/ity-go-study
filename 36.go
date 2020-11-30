/*
* @Author: 刘东旭
* @Date: 2020-11-22 12:23:14
 * @LastEditors: 刘东旭
 * @LastEditTime: 2020-11-30 18:51:59
* @Description:
*/
package main

import "fmt"

// 首字母大写 Person 表示公有结构体
// 首字母小写 person 表示私有结构体
type Person struct {
	name string
	age  int
	sex  string
}

func main() {
	var p1 Person //实例化Person结构体
	p1.name = "张三"
	p1.sex = "男"
	p1.age = 20
	// 值：{name:张三 age:20 sex:男} 类型: main.Person
	// fmt.Printf("值：%v 类型: %T\n", p1, p1)
	// 值：main.Person{name:"张三", age:20, sex:"男"} 类型: main.Person
	// fmt.Printf("值：%#v 类型: %T\n", p1, p1)

	var p2 = new(Person)
	p2.name = "李四"
	(*p2).sex = "女" //等同于上行
	p2.age = 38
	// 值：&main.Person{name:"李四", age:38, sex:"女"} 类型: *main.Person
	fmt.Printf("值：%#v 类型: %T\n", p2, p2)

}
