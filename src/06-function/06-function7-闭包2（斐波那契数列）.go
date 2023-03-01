package main

import "fmt"

// 实现斐波那契数列，利用闭包函数引用的变量具有记忆作用
// 1、1、2、3、5、8、13、21、34
func main() {
	//递归实现方式
	result := 0
	for i := 1; i <= 10; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}

	//闭包实现方式
	f := fibonacciV2()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

/**
 * 递归函数
 */
func fibonacci(n int) (res int) {
	if n <= 2 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

/**
 * 闭包实现方式
 */
func fibonacciV2() func() int {
	//这里的最上一层的func的返回值类型 是通过 最后一层的结果值一层层往上浮的，value的值类型是int，函数体的返回类型就是int，再往上一层返回类型也理应是int
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
