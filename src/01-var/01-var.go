package main

import (
	"errors"
	"fmt"
	"math"
)

// 声明全局变量
var (
	b3 int    = 1
	b4 string = "hb"
)

func main() {
	fmt.Println(errors.New("math: square root of negative number"))

	//1、最常规写法
	var name string = "huangbin"
	name = "huangbin123" //修改变量值
	fmt.Println(name)

	//2、省略数据类型，根据赋值自行判定变量类型
	var name1 = "huangbin1"
	fmt.Println(name1)

	//3、省略声明关键字
	//它只能被用在函数体内，而不可以用于全局变量的声明与赋值。使用操作符 := 可以高效地创建一个新的变量，称之为初始化声明。
	//如果变量已经被声明过，但依然使用:=时编译器会报错
	name2 := "huangbin2"
	fmt.Println(name2)

	//4、同时一行声明多个变量
	var e, f, g string = "1", "2", "3"
	fmt.Println(e, f, g)

	//5、同时声明不同数据类型
	var (
		a1 int  = 1
		b2 bool = true
	)

	fmt.Println(a1, b2, b3, b4)

	//6、匿名变量，任何赋给这个标识符的值都将被抛弃
	var _, b = 5, 7
	fmt.Println(b) //7

	/**
	变量的生命周期
	1、全局变量：它的生命周期和整个程序的运行周期是一致的；
	2、局部变量：它的生命周期则是动态的，从创建这个变量的声明语句开始，到这个变量不再被引用为止；
	3、形式参数和函数返回值：它们都属于局部变量，在函数被调用的时候创建，函数调用结束后被销毁。
	*/

	//在每次循环的开始会创建临时变量 i，然后在每次循环迭代中创建临时变量 x 和 y。
	//临时变量 x 存放在栈中，随着函数执行结束（ 执行遇到最后一个`}` ），释放其内存。
	//局部变量 sum 的生命周期则是动态的
	sum := 0
	for i := 0; i <= 3; i++ {
		x := math.Sin(30)
		fmt.Println(x)
		sum += i
	}
	fmt.Println(sum)

}
