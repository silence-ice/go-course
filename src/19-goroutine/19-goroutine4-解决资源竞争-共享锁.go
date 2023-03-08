package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter1 int64
	wg1      sync.WaitGroup
)

func main() {
	wg1.Add(2)
	go incCounter1()
	go incCounter1()

	wg1.Wait() //等待goroutine结束
	fmt.Println(counter1)
}

/*
*
共享锁
1、Go语言提供了传统的同步 goroutine 的机制，就是对共享资源加锁。atomic 包里的原子函数就可以对共享的资源进行加锁操作。
2、使用了 atmoic 包的 AddInt64 函数原理：

	3.1 这个函数会同步整型值的加法，方法是强制同一时刻只能有一个 gorountie 运行并完成这个加法操作。
	3.2 当 goroutine 试图去调用任何原子函数时，这些 goroutine 都会自动根据所引用的变量做同步处理。

3、另外两个有用的原子函数是 LoadInt64 和 StoreInt64。这两个函数提供了一种安全地读和写一个整型值的方式。
*/
func incCounter1() {
	defer wg1.Done()
	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter1, 1) //安全的对counter1加1
		runtime.Gosched()
	}
}
