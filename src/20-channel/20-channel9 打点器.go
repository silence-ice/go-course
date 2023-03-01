package main

import (
	"fmt"
	"time"
)

/*
*
打点器和计时器：
1、计时器（Timer）的原理和倒计时闹钟类似，都是给定多少时间后触发；返回 time.Ticker 对象。
2、打点器（Ticker）的原理和钟表类似，钟表每到整点就会触发；返回time.Timer 对象。
3、这两个对象里面可以通过一个 C 成员，类型是只能接收的时间通道（<-chan Time），使用这个通道就可以获得时间触发的通知。
*/
func main() {
	// 创建一个打点器, 每500毫秒触发一次
	ticker := time.NewTicker(time.Millisecond * 1000)

	// 声明计数变量
	var i int

	// 不断地检查通道情况
	for {
		fmt.Println("sdsd")
		// 多路复用通道
		select {
		case <-ticker.C: // 打点器触发了
			// 记录触发了多少次
			i++
			fmt.Println("tick", i)
		}
	}
}
