package main

import (
	"fmt"
	"time"
)

/*
*
什么是goroutine？
1、当使用者分配足够多的任务，系统能自动帮助使用者把任务分配到 CPU 上，让这些任务尽量并发运作，这种机制被称为 goroutine
2、goroutine 其实就是线程，但是它比线程更小，十几个 goroutine 可能体现在底层就是五六个线程。
3、Go 程序从 main 包的 main() 函数开始，在程序启动时，Go 程序就会为 main() 函数创建一个默认的 goroutine。
*/
func main() {
	//实现每隔一秒打印一次计数，而 main 的 goroutine 则等待用户输入，两个行为可以同时进行
	//使用 go 关键字创建 goroutine 时，被调用函数的返回值会被忽略，如果需要在 goroutine 中返回数据，需要通过channel把数据从 goroutine 中作为返回值传出
	go running()

	// 接受命令行输入, 不做任何事情
	var input string
	fmt.Scanln(&input)

	//所有 goroutine 在 main() 函数结束时会一同结束。
}

/**
【执行流程】该程序在运行过程中，总共会产生 2个 goroutine：
1、运行时（runtime）会默认为 main() 函数创建一个 goroutine。
2、在 main() 函数的 goroutine 中执行到 go running 语句时，归属于 running() 函数的 goroutine 被创建，running() 函数开始在自己的 goroutine 中执行。
3、此时，main() 继续执行，两个 goroutine 通过 Go 程序的调度机制同时运作
*/

func running() {
	times := 0
	for {
		fmt.Println(times)
		times++

		// 延时1秒
		time.Sleep(time.Second)
	}
}
