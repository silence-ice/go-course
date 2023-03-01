package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	// 创建一个无缓冲字符串通道
	ch := make(chan string)

	// 并发执行服务器逻辑
	go RPCServer(ch)

	// 客户端请求数据和接收数据
	result, err := RPCClient(ch, "hi")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

// 模拟RPC客户端的请求和接收消息封装
func RPCClient(ch chan string, req string) (string, error) {
	// 向服务器发送请求
	ch <- req

	// 等待服务器返回
	select {
	case data := <-ch: //读取ch
		return data, nil

	case <-time.After(1 * time.Second): //2秒后触发
		return "", errors.New("Time out")
	}
}

// 模拟RPC服务器端接收客户端请求和回应
func RPCServer(ch chan string) {
	for {
		// 接收客户端请求
		data := <-ch
		// 打印接收到的数据
		fmt.Println("server received:", data)

		// 模拟超时，通过睡眠函数让程序执行阻塞2秒的任务，走select的另外一个逻辑
		//time.Sleep(time.Second * 2)

		// 反馈给客户端收到
		ch <- "roger"
	}
}
