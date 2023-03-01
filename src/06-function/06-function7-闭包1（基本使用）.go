package main

import "fmt"

/**
什么是闭包？
* 通俗地讲就是"别人家"有某个东西，你想拿到但是因为权限不够（不打死你才怪），但是你可以跟家里的孩子套近乎，通过他拿到！
* 这个家就是"局部作用域"，外部无法访问内部变量，孩子是返回对象/函数，对家里的东西有访问权限，借助返回对象/函数间接访问内部变量！
*/

/**
闭包的用途？
* 一个是前面提到的可以读取函数内部的变量
* 另一个就是让这些变量的值始终保持在内存中，不会在 f1 调用后被自动清除。
*/

/*
*
1、Go语言中闭包是引用了自由变量的函数，闭包 = 函数 + 引用环境。
2、闭包可以对它作用域上部的变量进行引用，修改引用的变量也会进行实际修改.
3、被引用到闭包中的变量让闭包本身拥有了记忆效应，变量会跟随闭包生命期一直存在，闭包本身就如同变量一样拥有了记忆效应。
*/
func main() {
	//one
	a := 1

	f := func() int {
		//在匿名函数中并没有定义a，a的定义在匿名函数之前，此时a就被引用到了匿名函数中形成了闭包
		a++
		return a
	}

	fmt.Println(f())

	//two
	//对比输出的日志发现 kk 与 ff 输出的函数地址不同，因此它们是两个不同的闭包实例，所以两个闭包互不影响
	ff := Accumulate(1)     //创建一个累加器, 并返回一个闭包函数
	fmt.Println(ff())       //2
	fmt.Println(ff())       //3
	fmt.Printf("%p\n", &ff) //0xc00000e020，打印累加器的函数地址

	kk := Accumulate(2)
	fmt.Println(kk())       //3
	fmt.Println(kk())       //4
	fmt.Printf("%p\n", &kk) //0xc000080008

	//three
	//实现类似于设计模式中工厂模式的生成器，创意一个玩家生成器，并返回名字（name）和血量（hp）
	player := playerGen("silence")
	name, hp := player()
	fmt.Println(name, hp) //silence 150
}

/**
 * 累加器
 */
func Accumulate(value int) func() int {
	//这里的最上一层的func的返回值类型 是通过 最后一层的结果值一层层往上浮的，value的值类型是int，函数体的返回类型就是int，再往上一层返回类型也理应是int
	return func() int {
		value++
		return value
	}
}

/*
 * 玩家生成器
 */
func playerGen(name string) func() (string, int) {
	hp := 150

	return func() (string, int) {
		return name, hp
	}
}
