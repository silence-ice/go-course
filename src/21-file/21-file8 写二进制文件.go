package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

func main() {
	info := "http://c.biancheng.net/golang/"
	file, err := os.Create("./output.gob")
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(info)
	if err != nil {
		fmt.Println("编码错误", err.Error())
		return
	} else {
		fmt.Println("编码成功")
	}
}
