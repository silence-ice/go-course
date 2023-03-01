package main

import "fmt"

func main() {
	//one 在没有可变参数的情况下只能使用数组切片，也就是[]type
	fmt.Println(myfunc2([]int{1, 3, 7, 13}))
	//two go使用了类型...type代替了数组切片
	fmt.Println(sum(1, 2, 3))
	//three 使用interface{} 传递任意类型数据
	MyPrintf(1, "hb", 1.23)
}

func myfunc2(args []int) int {
	s := 0
	for _, arg := range args {
		s += arg
	}
	return s
}

/*
*
1、可变参数是指函数传入的参数个数是可变的
2、形如...type格式的类型只能作为函数的参数类型存在，并且必须是最后一个参数，它是一个语法糖（syntactic sugar）
*/
func sum(numbers ...int) int {
	s := 0
	for _, val := range numbers {
		s += val
	}
	return s
}

func MyPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}
