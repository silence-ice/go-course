package main

import (
	"fmt"
	"time"
)

/*
*
打点器和计时器：
1、计时器（Timer）的原理和倒计时闹钟类似，都是给定多少时间后触发；返回 time.Ticker 对象。
3、对象里面可以通过一个 C 成员，类型是只能接收的时间通道（<-chan Time），使用这个通道就可以获得时间触发的通知。
*/
func main() {
	// 创建一个计时器, 2秒后触发
	stopper := time.NewTimer(time.Second * 2)

	// 不断地检查通道情况
	for {
		// 多路复用通道
		select {
		case <-stopper.C: // 2秒后计时器到时了
			fmt.Println("重置定时器")
			stopper.Reset(time.Second * 2) //重置定时器
		}

		//2秒后输出（每次命中stopper.C才会走到这里，未到期select会阻塞住）
		fmt.Println("2秒后输出")
	}
}
