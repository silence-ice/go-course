package main

import (
	"fmt"
	"strings"
)

/*
*
1、批量处理，将下面字符串处理成下面字符串

	go scanner => SCANNER
	go parser => PARSER
	go compiler => COMPILER
	go printer => PRINTER
	go formater => FORMATER

2、mainBak是实验过程步骤
*/
func mainBak() {
	//one 测试单个字符串的处理过程
	s := "go hb"
	s = removePrefix(s)
	fmt.Println(s) // hb
	s = strings.TrimSpace(s)
	fmt.Println(s) //hb
	s = strings.ToUpper(s)
	fmt.Println(s) //HB

	//two 测试函数赋值
	var chain func(string) string
	chain = removePrefix
	fmt.Println(chain("go hb")) // hb

	//three 测试函数集合赋值
	//	var b = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	var chain1 = []func(string) string{
		removePrefix,
		strings.TrimSpace,
		strings.ToUpper,
	}
	//fmt.Println(chain1("go hb"))

	// 待处理的字符串列表
	list := []string{
		"go scanner",
		"go parser",
		"go compiler",
		"go printer",
		"go formater",
	}

	for _, val := range list {
		fmt.Println(val)
	}

	for index, val := range list {
		result := val
		//fmt.Println(result)
		for _, funcOne := range chain1 {
			result = funcOne(result)
		}

		list[index] = result
	}

	for _, val := range list {
		fmt.Println(val)
	}
}

func main() {
	//mainBak()

	// 待处理的字符串列表
	list := []string{
		"go scanner",
		"go parser",
		"go compiler",
		"go printer",
		"go formater",
	}

	// 处理函数链
	var chain = []func(string) string{
		removePrefix,
		strings.TrimSpace,
		strings.ToUpper,
	}

	// 进入处理方法
	result := StringProccess(list, chain)

	for _, val := range result {
		fmt.Println(val)
	}
}

func StringProccess(list []string, chain []func(string) string) []string {
	for index, val := range list {
		result := val
		//fmt.Println(result)
		for _, funcOne := range chain {
			result = funcOne(result)
		}

		list[index] = result
	}

	return list
}

/**
 * 自定义的移除前缀的处理函数
 */
func removePrefix(str string) string {
	return strings.TrimPrefix(str, "go")
}
