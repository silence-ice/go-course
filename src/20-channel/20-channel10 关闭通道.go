package main

import "fmt"

func main() {
	//关闭通道后的长度和容量
	ch := make(chan int) // 创建一个整型带两个缓冲的通道
	close(ch)            // 关闭通道

	v, ok := <-ch                 // 打印取出数据的状态
	fmt.Println(v, ok)            //0 false
	fmt.Println(cap(ch), len(ch)) //输出容量和长度，0 0

	//带缓冲通道的数据，表明缓冲通道在关闭后依然可以访问内部的数据
	ch1 := make(chan int, 2)

	ch1 <- 0
	ch1 <- 1
	close(ch1)

	// 遍历缓冲所有数据, 且多遍历个
	for i := 0; i < cap(ch1)+1; i++ {

		v, ok := <-ch1
		/**
		0 true
		1 true
		0 false
		*/
		fmt.Println(v, ok)
	}
}
