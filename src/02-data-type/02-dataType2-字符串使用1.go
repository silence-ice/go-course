package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func main() {
	//使用双引号""来定义字符串
	//字符串中可以使用转义字符来实现换行、缩进等效果，\n：换行符 <==>  \r：回车符 <==>  \t：tab 键 <==> \u 或 \U：Unicode 字符 <==> \\：反斜杠自身
	str1 := "C语言中文网\nGo语言教程"
	fmt.Println(str1)

	//字符串拼接符“+”
	str2 := "Beginning of the string " + "second part of the string"
	fmt.Println(str2)

	//字符串拼接
	//简单的东西未必高效。除了使用+连接字符串，Go 也有类似于 StringBuilder 的机制来进行高效的字符串连接
	hammer := "吃我一锤"
	sickle := "死吧"

	// 声明字节缓冲，bytes.Buffer 是可以缓冲并可以往里面写入各种字节数组的，字符串也是一种字节数组，
	var stringBuilder bytes.Buffer

	// 把字符串写入缓冲
	stringBuilder.WriteString(hammer)
	stringBuilder.WriteString(sickle)

	// 将缓冲以字符串形式输出
	fmt.Println(stringBuilder.String()) //吃我一锤死吧

	//定义多行字符串
	const str3 = `第一行
				第二行
				第三行
				\r\n
				`
	fmt.Println(str3)

	//计算字符串长度
	str4 := "黄斌"
	str5 := "hb"
	fmt.Println(len(str4))                    //输出6，Go 语言的字符串都以 UTF-8 格式保存，每个中文占用 3 个字节
	fmt.Println(len(str5))                    //输出2，ASCII 字符串长度可以直接使用 len() 函数
	fmt.Println(utf8.RuneCountInString(str4)) //输出2，Unicode 字符串（中文、日文）长度使用 utf8.RuneCountInString() 函数

	//字符串格式化输出
	var progress = 2
	var target = 8
	// 两参数格式化
	title := fmt.Sprintf("已采集%d个药草, 还需要%d个完成任务", progress, target)
	fmt.Println(title) //已采集2个药草, 还需要8个完成任务
	// 按数值本身的格式输出
	variant := fmt.Sprintf("%v %v %v", "月球基地", 3.14159, true) //月球基地 3.14159 true
	fmt.Println(variant)
	/**
	  %v	按值的本来值输出
	  %+v	在 %v 基础上，对结构体字段名和值进行展开
	  %#v	输出 Go 语言语法格式的值
	  %T	输出 Go 语言语法格式的类型和值
	  %%	输出 % 本体
	  %b	整型以二进制方式显示
	  %o	整型以八进制方式显示
	  %d	整型以十进制方式显示
	  %x	整型以十六进制方式显示
	  %X	整型以十六进制、字母大写方式显示
	  %U	Unicode 字符
	  %f	浮点数
	  %p	指针，十六进制方式显示
	*/
}
