package main

import (
	"fmt"
	"time"
)

/*
*
如何设置超时呢？
1、答案是使用 select 来设置超时。
2、select 的用法与 switch 语言非常类似，但select 有比较多的限制，其中最大的一条限制就是每个 case 语句里必须是一个 IO 操作

	select {
		case <-chan1:
			// 如果chan1成功读到数据，则进行该case处理语句
		case chan2 <- 1:
			// 如果成功向chan2写入数据，则进行该case处理语句
		default:
			// 如果上面都没有成功，则进入default处理流程
	}

3、没有命中case的情况：

	1、如果给出了 default 语句，那么就会执行 default 语句，同时程序的执行会从 select 语句后的语句中恢复；
	2、如果没有 default 语句，那么 select 语句将被"阻塞"，直到至少有一个通信可以进行下去
*/
func main() {
	//两个goruntine使用2个通道（ch、quit）
	ch := make(chan int)
	quit := make(chan bool)

	//新开一个协程
	go func() {
		for {
			select {
			case num := <-ch: //读取ch
				fmt.Println("num = ", num)

			case <-time.After(3 * time.Second): //3秒后触发
				fmt.Println("超时")
				quit <- true
			}
		}
	}()

	for i := 0; i < 5; i++ {
		ch <- i
		//time.Sleep(time.Second)//睡眠一秒，可以不需要
	}

	<-quit //需要有这个，不然当i=4的时候，main的goruntine可能直接退出了，等待quit通道
	fmt.Println("程序结束")

	/**
	输出内容：
		num =  0
		num =  1
		num =  2
		num =  3
		num =  4
		超时
		程序结束
	*/
}
