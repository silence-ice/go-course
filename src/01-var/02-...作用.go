package main

func main() {
	newSum(1, 2, 3)

	//第二个用法是slice可以被打散进行传递。
	var numList = []int{1, 2, 3}
	newSum(numList...)
}

func newSum(num ...int) {
	//第一个用法主要是用于函数有多个不定参数的情况，可以接受多个不确定数量的参数。
	for _, val := range num {
		println(val)
	}
}
