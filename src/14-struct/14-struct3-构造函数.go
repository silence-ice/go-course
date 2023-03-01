package main

import "fmt"

type Cat struct {
	Color string
	Name  string
}

func main() {
	/**
	1、Go语言中没有提供构造函数相关的特殊机制，可以根据自己的需求将参数使用函数传递到结构体，并返回结构体即可完成构造函数的任务
	*/
	fmt.Println(initCatName("silence"))        //&{ silence}
	fmt.Println(initCatColor("blue"))          //&{blue }
	fmt.Println(initAllCat("silence", "blue")) //&{silence blue}
}

// 定义Name属性构造函数
func initCatName(name string) *Cat {
	return &Cat{
		Name: name,
	}
}

// 定义Color属性构造函数
func initCatColor(color string) *Cat {
	return &Cat{
		Color: color,
	}
}

// 定义Name、Color属性构造函数
func initAllCat(color string, name string) *Cat {
	return &Cat{
		Color: color,
		Name:  name,
	}
}
