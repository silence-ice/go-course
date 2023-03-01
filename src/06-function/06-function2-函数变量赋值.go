package main

import "fmt"

/**
 * 函数也是一种类型，可以和其他类型一样保存在变量中
 */
func main() {
	var t func() //将变量 t 声明为 func() 类型，此时 t 就被俗称为“回调函数”，此时 t 的值为 nil
	t = test     //将 test() 函数作为值，赋给函数变量 t
	t()

	var t2 func(string) string
	t2 = test2
	fmt.Println(t2("test2"))

}

func test() {
	fmt.Println("test")
}

func test2(str string) string {
	return str
}
