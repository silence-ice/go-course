package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	/**
	1、bufio.NewWriter()和bytes.Buffer的 Write()功能是一样的
	2、可以到各自的方法下面查看，发现携带的方法有各自不同，比如Flush()、Buffered()、ReadString()是bufio才有的
	*/
	//写入方式一 => bytes.Buffer 写入
	var writer bytes.Buffer
	n, err := writer.Write([]byte("Huang Bin")) //字节类型写入
	fmt.Println(writer.String(), n, err)        //Huang Bin 9 <nil>

	//写入方式二 => bufio.NewWriter 写入
	wr := bytes.NewBuffer([]byte("")) //等同bytes.NewBuffer(nil)
	newWriter := bufio.NewWriter(wr)

	n1, err1 := newWriter.Write([]byte("Huang Bin")) //字节类型写入
	newWriter.Flush()

	fmt.Println(wr.String(), n1, err1)        //Huang Bin 9 <nil>
	fmt.Println(string(wr.Bytes()), n1, err1) //Huang Bin 9 <nil>

	////////////////bufio.NewWriter()常用方法////////////////
	/**
	1、Available() 方法的功能是返回缓冲区中未使用的字节数
		func (b *Writer) Available() int
	2、Buffered() 方法的功能是返回已写入当前缓冲区中的字节数
		func (b *Writer) Buffered() int
	3、Flush() 方法的功能是把缓冲区中的数据写入底层的 io.Writer，并返回错误信息。如果成功写入，error 返回 nil，否则 error 返回错误原因：
		func (b *Writer) Flush() error
	*/
	wr1 := bytes.NewBuffer([]byte(""))
	newWriter1 := bufio.NewWriter(wr1)

	fmt.Println("写入前已使用的缓冲区为：", newWriter1.Buffered()) //0
	fmt.Println("写入前可用的缓冲区为：", newWriter1.Available()) //4096

	n2, err2 := newWriter1.Write([]byte("黄斌"))

	fmt.Printf("写入后，可用的缓冲区为：%d\n", newWriter1.Available()) //4090
	fmt.Printf("写入后，已使用的缓冲区为：%d\n", newWriter1.Buffered()) //6

	err3 := newWriter1.Flush()

	fmt.Printf("写入后，已使用的缓冲区为：%d\n", newWriter1.Buffered()) //0

	fmt.Println(string(wr1.Bytes()), n2, err2, err3) //黄斌 6 <nil> <nil>
}
