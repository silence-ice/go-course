package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	/**
	什么是读取器？
	1、io.Reader 表示一个读取器，它将数据从 "某个资源" 读取到 "传输缓冲区"。在缓冲区中，数据可以被流式传输和使用。
	2、只要实现了  Read(p []byte) 方法，那它就是一个读取器。
		type Reader interface {
			Read(p []byte) (n int, err error)
		}

		2.1 Read() 方法的功能是读取数据，并存放到字节切片 p 中。
		2.2 Read() 执行结束会返回已读取的字节数 n，因为最多只调用底层的 io.Reader 一次，所以返回的 n 可能小于 len(p)，
		2.3 当字节流结束时，n 为 0，err 为 io. EOF。
	3、strings包提供了NewReader()方法，可以看到里面的struct实现了 Read(p []byte) 方法。
	*/
	reader := strings.NewReader("C语言中文网") //reader实际就是个io.Reader
	//fmt.Println(reader)//&{C语言中文网 0 -1}

	/**
	buf := make([]byte, 4)
	n, err := reader.Read(b)
	fmt.Println(string(b[:n]), n, err)//C语 4 <nil>
	功能同下
	*/
	var buf [4]byte
	n, err := reader.Read(buf[:])
	fmt.Println(string(buf[:n]), n, err) //C语 4 <nil>

	n1, err1 := reader.Read(buf[:])
	fmt.Println(string(buf[:n]), n1, err1) //言� 4 <nil>

	for {
		n, err := reader.Read(buf[:])
		if err != nil && err == io.EOF {
			fmt.Println("EOF:", n, err) //EOF: 0 EOF
			break
		}
		fmt.Println(string(buf[:n]), n1, err1) //��� 4 <nil>  �网 4 <nil>
	}

	//////////////////////////////////////////////////////////////////////////////////////////////////////
	/**
	什么是编写器？
	1、io.Writer 表示一个编写器，它从缓冲区读取数据，并将数据写入目标资源。
	2、只要实现了 Write(p []byte) ，那它就是一个编写器。
		type Writer interface {
			Write(p []byte) (n int, err error)
		}


		2.1 Write() 方法的功能是写入数据，字节切片 p写入目标资源。
		2.2 Write() 执行结束会返回已写入的字节数 n
		2.3 当字节流结束时，n 为 0，err 为 io. EOF。
	3、bytes包提供了Buffer struct，可以看到里面的struct实现了 Write(p []byte) 方法。
	*/
	stringList := []string{
		"Hello",
		"World",
		"Huang",
		"Bin",
	}

	//使用 bytes.Buffer 类型作为 io.Writer 将数据写入内存缓冲区
	var writer bytes.Buffer

	for _, val := range stringList {
		//将val数据写入writer变量中，其中[]byte(val)表示将utf-8转byte
		n, err := writer.Write([]byte(val))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if n != len(val) {
			fmt.Println("failed to write data")
			os.Exit(1)
		}
	}

	fmt.Println(writer.String()) //HelloWorldHuangBin
}
