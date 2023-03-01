package main

import "fmt"

/**
什么是channel？
1、channel 可以让一个 goroutine 通过它给另一个 goroutine 发送值信息。
3、在任何时候，同时只能有一个 goroutine 访问 channel 进行发送和获取数据。goroutine 间通过 channel 就可以通信。
3、channel 像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。

声明通道类型
1、声明格式：var 通道变量 chan 通道类型
2、每个 channel 都有一个可发送数据的类型。一个可以发送 int 类型数据的 channel 一般写为 chan int。

创建通道
1、通道实例 := make(chan 数据类型)
2、创建示例
	ch1 := make(chan int)                 // 创建一个整型类型的通道
	ch2 := make(chan interface{})         // 创建一个空接口类型的通道, 可以存放任意格式

	type Equip struct{ }
	ch2 := make(chan *Equip)             // 创建Equip指针类型的通道, 可以存放*Equip

接收通道数据
1、通道的收发操作在不同的两个 goroutine 间进行。
2、如果通道的数据在没有接收方处理时，数据发送方会持续阻塞，因此通道的接收必定在另外一个 goroutine 中进行。
3、通道一次只能接收一个数据元素。

1) 阻塞接收数据：data := <-ch
2) 接收任意数据，忽略接收的数据：<-ch

*/

// 限循环接收打印，当遇到数据0时, 退出接收循环
func main() {
	//构建一个通道
	ch1 := make(chan int)

	wg.Add(1)
	// 开启一个并发匿名函数
	go func() {
		for i := 3; i >= 0; i-- {
			// 通过channel通知main的goroutine
			ch1 <- i
		}
		wg.Done()
	}()

	//等待goroutine结束，不能在这行写 wg.Wait()
	//因为匿名函数的goroutine在通道的数据在没有接收方处理时，数据发送方会持续阻塞，会出现死循环
	//wg.Wait()

	for val := range ch1 {
		fmt.Println(val)

		// 当遇到数据0时, 退出接收循环，也可以在匿名函数使用 close(ch) 来关闭通道，否则range会一直循环下去阻塞住
		if val == 0 {
			break
		}
	}

	wg.Wait() //等待goroutine结束，这行不能移上

	//如果出现下面错误： fatal error: all goroutines are asleep - deadlock!
	//报错的意思是：运行时发现所有的 goroutine（包括main）都处于等待 goroutine。也就是说所有 goroutine 中的 channel 并没有形成发送和接收对应的代码。简而言之在 同一时刻 只有发送，没有接受就会报错！！
}
