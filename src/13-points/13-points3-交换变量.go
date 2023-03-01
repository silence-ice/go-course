package main

import "fmt"

func main() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200

	failSwap(a, b) //100 200
	fmt.Println(a, b)

	swap(&a, &b)
	fmt.Println(a, b)
}

/**
 * 无法交换值，因为go语言的变量交换是值传递，不像php一样是引用传递
 */
func failSwap(a int, b int) {
	temp := a
	a = b
	b = temp
}

/**
 * 正确交换值的方式
 */
func swap(x *int, y *int) {
	var temp int
	temp = *x /* 保存 x 地址的值 */
	*x = *y   /* 将 y 赋值给 x */
	*y = temp /* 将 temp 赋值给 y */

	//*x, *y = *y, *x//简洁写法
}
