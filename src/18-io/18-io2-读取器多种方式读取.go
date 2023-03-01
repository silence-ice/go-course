package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	//普通读取
	data, _ := ReadFrom(strings.NewReader("中华民族是个伟大的民族"), 80)
	fmt.Println(string(data)) //中华民族是个伟大的民族

	//文件读取
	file, _ := os.Open("./read.txt")
	data2, _ := ReadFrom(file, 80)
	fmt.Println(string(data2)) //中华民族是个伟大的民族

	//命令行读取
	data3, _ := ReadFrom(os.Stdin, 80)
	fmt.Println(string(data3)) //输入什么就输出什么
}

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)

	if n > 0 {
		return p[:n], nil
	}

	return p, err
}
