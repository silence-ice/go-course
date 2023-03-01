package main

import "fmt"

// 使用生产者和消费者，无限循环接收打印，当遇到数据0时, 退出接收循环
func main() {
	/**
	1、使用无缓冲通道往里面装入数据的时候，装入方将被阻塞
	2、如果通道中没有放入任何数据，接收方试图从通道中获取数据时，同样也是阻塞
	*/
	ch := make(chan int)

	//这行代码不能移到 close(ch)上面，因为同一时刻 要有发送，也要有接受，不然就会报错！！
	go printer(ch)

	for i := 1; i <= 10; i++ {
		// 将数据通过channel投送给printer
		ch <- i
	}

	// 通知并发的printer结束循环(没数据啦!)
	ch <- 0

	//关闭通道
	close(ch)

	//判断channel是否被关闭
	//（阻塞）接收数据：data := <-ch
	_, ok := <-ch

	if !ok {
		fmt.Println("关闭通道")
	}
}

func printer(ch chan int) {
	// 开始无限循环等待数据
	for {
		// 从channel中获取一个数据
		data := <-ch
		fmt.Println(data)

		// 将0视为数据结束
		if data == 0 {
			break
		}
	}
}
