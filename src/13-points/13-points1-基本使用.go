package main

import "fmt"

func main() {
	/**
	指针的作用：
		1、指针主要是防止变量的拷贝导致资源的消耗

	指针的特性和使用方法：
		1、对变量进行取地址操作使用&操作符，可以获得这个变量的指针变量。
		2、指针变量的值是指针地址。
		3、对指针变量进行取值操作使用*操作符，可以获得指针变量指向的原变量的值。
	*/
	//fmt.Println(time.Now())

	var a int = 20 /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */

	//变量 是一种使用方便的占位符，用于引用计算机内存地址
	fmt.Println("变量的值:", a)       //20
	fmt.Printf("变量的存储地址:", &a)    //%!(EXTRA *int=0xc000086000)
	fmt.Printf("变量的存储地址值:%x", &a) //c000086000

	//指针变量 指向了一个值的内存地址，指针声明格式如下：var var_name *var-type
	ip = &a //指定指针变量的存储地址

	fmt.Printf("a 变量的地址是: %x\n", &a)     //输出变量的存储地址 c000096000
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip) //输出指针变量的存储地址 c000096000
	fmt.Printf("*ip 变量的值: %d\n", *ip)    //使用指针访问变量值 20

	//当一个指针被定义后没有分配到任何变量时，它的值为 nil。
	//nil 指针也称为空指针。
	//nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。
	var ptr *int
	fmt.Printf("ptr 的值为 : %x\n", ptr) //0

	if ptr == nil {
		fmt.Println("ptr 是空指针")
	}
}
