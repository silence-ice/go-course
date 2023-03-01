package main

import (
	"fmt"
	"time"
)

/*
*
time.AfterFunc() 函数
1、time.AfterFunc() 函数是在 time.After 基础上增加了到时的回调，方便使用。
2、time.After() 函数又是在 time.NewTimer() 函数上进行的封装
*/
func main() {
	// 声明一个退出用的通道，因为time.AfterFunc()方法里面的匿名函数变量是在另外一个goroutine中调用的
	exit := make(chan int)

	// 打印开始
	fmt.Println("start")

	//调用 time.AfterFunc() 函数，传入等待的时间和一个回调。
	//回调使用一个匿名函数，匿名函数会在另外一个 goroutine 中被调用，那么匿名函数处理的数据需要交到main的goroutine中，所以才需要使用channe
	time.AfterFunc(time.Second, func() {

		// 1秒后, 打印结果
		fmt.Println("one second after")

		// 通知main()的goroutine已经结束
		exit <- 0
	})

	// 等待结束
	data := <-exit
	fmt.Println(data)
}
