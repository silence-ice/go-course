package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

func main() {
	/**
	1、io.NewReader()和strings.NewReader的Read()功能是一样的
	2、可以到各自的方法下面查看，发现携带的方法有各自不同，比如Peek()、Buffered()、ReadString()是bufio才有的
	*/
	//读取方式一 => strings.NewReader读取
	reader := strings.NewReader("Hello World")

	by := make([]byte, 80)
	n, err := reader.Read(by)
	fmt.Println(string(by[0:n]), n, err) ///Hello World 11 <nil>

	//读取方式二 => bufio.NewReader读取
	reader1 := strings.NewReader("Hello World")
	bufioReader1 := bufio.NewReader(reader1)

	var buf [80]byte
	n1, err1 := bufioReader1.Read(buf[:])
	fmt.Println(string(buf[:n]), n1, err1) //Hello World 11 <nil>

	//读取方式三 => bytes.NewReader读取
	bytesReader := bytes.NewReader([]byte("Hello World")) //必须要转化成byte类型
	bufioReader2 := bufio.NewReader(bytesReader)

	by2 := make([]byte, 80)
	n2, err2 := bufioReader2.Read(by2)
	fmt.Println(string(by2[:n2]), n2, err2) //Hello World 11 <nil>

	/**
	Peek() 方法
	1、读取指定字节数的数据，这些被读取的数据不会从缓冲区中清除。在下次读取之后，本次返回的字节切片会失效。
	2、如果 Peek 返回的字节数不足 n 字节，则会同时返回一个错误说明原因，如果 n 比缓冲区要大，则错误为 ErrBufferFull。该方法原型如下：
		func (b *Reader) Peek(n int) ([]byte, error)
	*/
	stringReader := strings.NewReader("Huang And Ling")

	bufReader := bufio.NewReader(stringReader)
	peekData, _ := bufReader.Peek(3)
	fmt.Println(string(peekData)) //Hua

	peekData2, _ := bufReader.Peek(6)
	fmt.Println(string(peekData2)) //Huang

	//Buffered() 方法的功能是返回可从缓冲区读出数据的字节数
	fmt.Println(bufReader.Buffered()) //14

	/**
	ReadString()方法
	1、ReadString() 方法的功能是读取数据直到分隔符“delim”第一次出现，并返回一个包含“delim”的字符串。
	2、如果 ReadString 在读取到“delim”前遇到错误，它返回已读字符串和那个错误（通常是 io.EOF）
	*/
	var delim byte = ' '
	readStr, _ := bufReader.ReadString(delim)
	fmt.Println(readStr) //输出`Huang `

	//因为`Huang `从缓冲区读取出来了，只剩`And Ling`，所以缓存的就只剩8个字节数
	fmt.Println(bufReader.Buffered()) //8 = 14 - 6
}
