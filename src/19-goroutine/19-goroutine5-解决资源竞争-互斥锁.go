package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int64
	wg2     sync.WaitGroup
	mutex   sync.Mutex
)

func main() {
	wg2.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg2.Wait()
	fmt.Println(counter)
}

// 另一种同步访问共享资源的方式是使用互斥锁，互斥锁用于在代码上创建一个临界区，保证同一时间只有一个 goroutine 可以执行这个临界代码。
func incCounter(id int) {
	defer wg2.Done()

	for count := 0; count < 2; count++ {
		//同一时刻只有一个 goroutine 可以进入临界区。之后直到调用 Unlock 函数之后，其他 goroutine 才能进去临界区。
		mutex.Lock()
		{
			value := counter
			//当调用 runtime.Gosched 函数强制将当前 goroutine 退出当前线程后，调度器会再次分配这个 goroutine 继续运行。
			runtime.Gosched()
			value++
			counter = value
			//counter++
		}
		mutex.Unlock() //释放锁，允许其他正在等待的goroutine进入临界区
	}
}
