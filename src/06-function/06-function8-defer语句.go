package main

import (
	"fmt"
	"os"
)

/*
*
1、在 defer 归属的函数即将结束时（即函数返回return之前），defer 语句会将其后面跟随的语句进行延迟处理。
2、将延迟处理的语句按 defer 的逆序进行执行，也就是说，先被 defer 的语句最后被执行，最后被 defer 的语句，最先被执行。
*/
func main() {
	test1()                                          //输出three、two、one
	fmt.Println(fileSize1("./06-function4-匿名函数.go")) //输出1417
	fmt.Println(fileSize2("./06-function4-匿名函数.go")) //输出1417
}

func test1() {
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
}

// 根据文件名查询其大小
func fileSize1(filename string) int64 {
	// 根据文件名打开文件, 返回文件句柄和错误
	f, err := os.Open(filename)
	// 如果打开时发生错误, 返回文件大小为0
	if err != nil {
		return 0
	}
	// 取文件状态信息
	info, err := f.Stat()

	// 如果获取信息时发生错误, 关闭文件并返回文件大小为0
	if err != nil {
		f.Close()
		return 0
	}
	// 取文件大小
	size := info.Size()
	// 关闭文件
	f.Close()

	// 返回文件大小
	return size
}

// 根据文件名查询其大小（defer升级版）
func fileSize2(filename string) int64 {
	// 根据文件名打开文件, 返回文件句柄和错误
	f, err := os.Open(filename)
	// 如果打开时发生错误, 返回文件大小为0
	if err != nil {
		return 0
	}
	// 取文件状态信息
	info, err := f.Stat()

	//延迟调用Close, 此时Close不会被调用
	defer f.Close()

	//如果获取信息时发生错误, 关闭文件并返回文件大小为0
	if err != nil {
		return 0
	}
	// 取文件大小
	size := info.Size()

	// 返回文件大小
	return size
}
