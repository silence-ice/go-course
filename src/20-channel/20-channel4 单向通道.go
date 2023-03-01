package main

import "fmt"

/*
*
什么是单向通道？
1、单向 channel 只能用于发送或者接收数据。channel 本身必然是同时支持读写的，否则根本没法用。
2、如果一个 channel 只允许写，即使写进去了，也没有丝毫意义，因为没有机会读取里面的数据。所谓的单向 channel 概念，其实只是对 channel 的一种使用限制。

单向通道声明方式：
1、var 通道实例 chan<- 元素类型    // 只能发送通道
2、var 通道实例 <-chan 元素类型    // 只能接收通道
*/
func main() {
	ch := make(chan int)
	// 声明一个只能发送的通道类型, 并赋值为ch
	var chSendOnly chan<- int = ch
	//声明一个只能接收的通道类型, 并赋值为ch
	var chRecvOnly <-chan int = ch

	fmt.Println(chSendOnly, chRecvOnly) //0xc00008c000 0xc00008c000
}
