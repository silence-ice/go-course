package main

import "fmt"

/*
*
1、recover可以让进入宕机流程中的 goroutine 恢复过来，如果报错程序就会调用 recover 可以捕获到 panic 的输入值，并且恢复正常的执行
2、在正常的执行过程中，调用 recover 会返回 nil 并且没有其他任何效果。
3、recover 仅在延迟函数 defer 中有效
*/
func main() {
	tryRecover()
}

func tryRecover() {
	//one recover 仅在延迟函数 defer 中有效
	//注册捕获panic的函数，必须先注册，若在panic之后则无效
	defer func() {
		r := recover()
		err, ok := r.(error) //获取recover的错误
		if ok {
			fmt.Println("Error occurred:", err)
		} else {
			panic(fmt.Sprintf("I don't know what to do : %v", r))
		}
	}()

	/**
	Error occurred: runtime error: integer divide by zero
	*/
	b := 0
	a := 5 / b
	fmt.Println(a)

	/**
	5
	panic: I don't know what to do : <nil>

	goroutine 1 [running]:
	main.tryRecover.func1()
	        /Users/silence/Documents/go/course/src/17-error/17-recover.go:22 +0xfa
	main.tryRecover()
	        /Users/silence/Documents/go/course/src/17-error/17-recover.go:30 +0xa2
	main.main()
	        /Users/silence/Documents/go/course/src/17-error/17-recover.go:11 +0x20
	exit status 2
	*/
	//b := 1
	//a := 5/b
	//fmt.Println(a)//5
}
