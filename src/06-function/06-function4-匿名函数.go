package main

import "fmt"

/*
*
匿名函数的定义格式如下：

	func(参数列表)(返回参数列表){
	    函数体
	}
*/
func main() {
	//one 匿名函数可以在声明后调用
	func(str string) string {
		fmt.Println(str)
		return str
	}("huangbin")

	//two 将匿名函数赋值给变量
	f := func(str string) string {
		fmt.Println(str)
		return str
	}
	f("huangbin")

	//three 匿名函数用作回调函数
	var list = []string{"huangbin", "liuhuiling"}
	result := visitList(list, func(s string) string {
		return s + "love"
	})
	fmt.Println(result)

	//four 使用map 值，通过key参数动态调用匿名函数
	var skill = map[string]func(){
		"fire": func() {
			fmt.Println("chicken fire")
		},
		"run": func() {
			fmt.Println("soldier run")
		},
		"fly": func() {
			fmt.Println("angel fly")
		},
	}

	fmt.Println(skill["run"]) //0x1093970
	if f, ok := skill["run"]; ok {
		f() //soldier run
	} else {
		fmt.Println("skill not found")
	}
}

/**
 *	数据处理
 *	1、遍历中访问每个元素的，然后对每个子元素使用匿名函数来处理
 *	2、第一个变量参数是字符串集合
 *	3、第二个变量参数：callback func(string) (string)
 *	4、第三个变量返回参数也是字符串集合
 */
func visitList(list []string, callback func(string) string) []string {
	result := list
	for key, val := range list {
		result[key] = callback(val)
	}

	return result
}
