package main

import (
	"fmt"
	"reflect"
)

func main() {
	//[整型]
	var num int
	fmt.Println(num) //默认为零值
	num++
	fmt.Println(num) //1
	//a = a++ // Go 的自增，自减只能作为表达式使用，会出现变异错误 syntax error: unexpected ++ at end of statement

	//[布尔]
	var c bool
	fmt.Println(c) //默认为 false

	//[字符串]
	var s string
	fmt.Println(s) //默认是空

	//数据类型转型
	var d4 = 2
	c3 := float32(d4)
	fmt.Println(c3)                 //2
	fmt.Println(reflect.TypeOf(c3)) //float32
}
