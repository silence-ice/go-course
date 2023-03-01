package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	//普通写入
	bytesBuff := bytes.NewBuffer([]byte(""))
	ioWriter := bufio.NewWriter(bytesBuff)

	n, err := ioWriter.Write([]byte("黄斌"))
	err2 := ioWriter.Flush()

	fmt.Println(string(bytesBuff.Bytes()), n, err, err2) //黄斌 6 <nil> <nil>

	//文件写入
	proverbs := []string{
		"Hello\n",
		"World\n",
	}
	file, err := os.Create("./write.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	for _, p := range proverbs {
		// 因为 os.Stdout 也实现了 io.Writer
		n, err := file.Write([]byte(p))
		if err != nil {
			fmt.Println(n, err)
			os.Exit(1)
		}
	}

	//命令行输出
	for _, p := range proverbs {
		// 因为 os.Stdout 也实现了 io.Writer
		n, err := os.Stdout.Write([]byte(p))
		if err != nil {
			fmt.Println(n, err)
			os.Exit(1)
		}
	}
}
