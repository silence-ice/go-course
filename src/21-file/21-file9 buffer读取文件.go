package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//printAllFile("01-var.go")
	printReadLineFile("01-var.go")
}

func printReadLineFile(filename string) {
	//读取所有文件内容
	file, err := os.Open(filename)

	if err != nil {
		//1、Go语言追求简洁优雅，所以，Go语言不支持传统的 try…catch…finally 这种异常。
		//2、在Go语言中，使用多值返回来返回错误。不要用异常代替错误，更不要用来控制流程，Go没有异常机制，但有panic/recover模式来处理错误
		//3、Panic可以在任何地方引发，一般会导致程序挂掉；但recover只有在defer调用的函数中有效
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text()) //打印每一行信息
	}
}

func printAllFile(filename string) {
	//读取所有文件内容
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(content)
	}
}
