package main

import (
	"fmt"
	"sync"
)

func main() {
	/**
	1、下面逻辑使用goroutine并发执行，导致所有 goroutine 在 main() 函数结束时会 一同结束，就会出现无法输出到出1～100的情况
	2、所以不得不使用time.Sleep来进行控制，但是对于实际生活的大多数场景来说，1秒是不够的，并且大部分时候我们都无法预知for循环内代码运行时间的长短。
	*/
	//	for i := 0; i < 100 ; i++{
	//		go fmt.Println(i)
	//	}
	////	time.Sleep(time.Second)//睡眠1秒

	/**
	WaitGroup的使用
	1、go语言中有一个其他的工具sync.WaitGroup 能更加方便的帮助我们达到这个目的。
	2、WaitGroup 对象内部有一个计数器，最初从0开始，它有三个方法：Add(), Done(), Wait() 用来控制计数器的数量。
		2.1 Add(n) 把计数器设置为n
		2.2 Done() 每次把计数器-1
		2.3 wait() 会阻塞代码的运行，直到计数器地值减为0。
	3、这里首先把wg计数设置为100， 每个for循环运行完毕都把计数器减一，主函数中使用Wait() 一直阻塞，直到wg为零——也就是所有的100个for循环都运行完毕
	*/
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()

}
