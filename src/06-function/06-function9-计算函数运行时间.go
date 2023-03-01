package main

import (
	"fmt"
	"time"
)

func main() {
	testTime()
}

func testTime() {
	start := time.Now() //获取当前时间 2020-01-30 13:25:52.33701 +0800 CST m=+0.000241845

	sum := 0
	for i := 0; i < 100000000; i++ {
		sum++
	}

	elapsed := time.Since(start) //Since() 函数表示返回从 t 到现在经过的时间，等价于time.Now().Sub(t)。
	//elapsed := time.Now().Sub(start)// time.Now().Sub() 的功能类似于 Since() 函数

	fmt.Println("该函数执行完成耗时：", elapsed)
}
