package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./output.gob")
	if err != nil {
		fmt.Println("文件打开失败", err.Error())
		return
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	info := ""
	err = decoder.Decode(&info)
	if err != nil {
		fmt.Println("解码失败", err.Error())
	} else {
		fmt.Println("解码成功")
		fmt.Println(info)
	}
}
