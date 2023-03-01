package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/**
案例场景：
1、在网球比赛中，两位选手会把球在两个人之间来回传递。选手总是处在以下两种状态之一，要么在等待接球，要么将球打向对方。
2、可以使用两个 goroutine 来模拟网球比赛，并使用无缓冲的通道来模拟球的来回。
3、这个示例程序展示2 个goroutine 间的网球比赛
*/

// wg 用来等待程序结束
var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

// main 是所有Go 程序的入口
func main() {
	// 创建一个无缓冲的通道
	court := make(chan int)

	// 计数加 2，表示要等待两个goroutine
	wg.Add(2)

	// 启动两个选手
	go player("黄斌", court)
	go player("刘慧玲", court)

	// 发球开始触发
	court <- 1

	// 等待游戏结束
	wg.Wait()
}

// player 模拟一个选手在打网球
func player(name string, court chan int) {
	// 在函数退出时调用Done 来通知main 函数工作已经完成
	defer wg.Done()

	for {
		// 等待球被击打过来（阻塞）
		ball, ok := <-court
		if !ok {
			// 如果通道被关闭，我们就赢了
			fmt.Printf("Player %s Won\n", name)
			return
		}

		// 选随机数，然后用这个数来判断我们是否丢球
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)

			// 关闭通道，表示我们输了
			close(court)
			return
		}

		// 显示击球数，并将击球数加1
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		// 将球打向对手
		court <- ball
	}
}
