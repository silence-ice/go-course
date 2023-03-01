package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Website1 struct {
	Name   string `xml:"name,attr"`
	Url    string
	Course []string
}

func main() {
	filePtr, err := os.Open("./info.json")
	if err != nil {
		fmt.Println("文件打开失败 [Err:%s]", err.Error())
		return
	}
	defer filePtr.Close()

	var info []Website1
	decoder := json.NewDecoder(filePtr) // 创建json解码器
	err = decoder.Decode(&info)

	if err != nil {
		fmt.Println("解码失败", err.Error())
	} else {
		fmt.Println("解码成功")
		fmt.Println(info)
	}
}
