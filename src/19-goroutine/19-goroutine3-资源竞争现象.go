package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	count int32
	wg    sync.WaitGroup
)

func main() {
	/**
	什么是资源竞争？
	1、如果两个或者多个 goroutine 在没有相互同步的情况下，访问某个共享的资源，比如同时对该资源进行读写时，就会处于相互竞争的状态。
	2、如果将程序多运行几次，会发现结果可能是 2、3、4。
	3、这是因为 count 变量没有任何同步保护，所以两个 goroutine 都会对其进行读写，会导致对已经计算好的结果被覆盖，以至于产生错误结果。
	*/
	wg.Add(2)
	go incCount()
	go incCount()
	wg.Wait()
	fmt.Println(count)

	/**
	下面我们来分析一下程序的运行过程，将两个 goroutine 分别假设为 g1 和 g2：
	1、g1 读取到 count 的值为 0；
	2、然后 g1 暂停了，切换到 g2 运行，g2 读取到 count 的值也为 0；
	3、g2 暂停，切换到 g1，g1 对 count+1，count 的值变为 1；
	4、g1 暂停，切换到 g2，g2 刚刚已经获取到值 0，对其 +1，最后赋值给 count，其结果还是 1；
	5、可以看出 g1 对 count+1 的结果被 g2 给覆盖了，两个 goroutine 都 +1 而结果还是 1。
	*/

	//共享资源竞争的问题，非常复杂，并且难以察觉，好在 Go 为我们提供了一个工具帮助我们检查，这个就是go build -race 命令，运行生成的可执行文件点击查看结果
}

func incCount() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		value := count
		//runtime.Gosched() 是让当前 goroutine 暂停的意思，退回执行队列，让其他等待的 goroutine 运行，目的是为了使资源竞争的结果更明显。
		runtime.Gosched()
		value++
		count = value
	}
}
