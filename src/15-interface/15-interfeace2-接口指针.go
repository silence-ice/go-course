package main

import (
	"fmt"
)

// //////////////////////////////接口方//////////////////////////////////
type Writer interface {
	Write(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

// //////////////////////////////实现方//////////////////////////////////
// 定义 Socket 结构
type Socket struct{}

// 要使用 Writer、Closer 接口，就必须实现接口的Write、Close方法，方法中的名称、参数列表、返回参数列表都必须一样
// 指针引用方式
func (s *Socket) Write(p []byte) (n int, err error) {
	return 0, nil
}

func (s *Socket) Close() error {
	return nil
}

func main() {
	var s Writer
	s = &Socket{}
	fmt.Println(s.Write(nil))

	var c Closer
	c = &Socket{}
	c.Close()

	//s := &Socket{}
	//fmt.Println(s.Write(nil))
	//s.Close()
}
