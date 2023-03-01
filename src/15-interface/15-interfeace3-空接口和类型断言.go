package main

import "fmt"

// 定义飞行动物接口
type Flyer1 interface {
	Fly()
}

// 定义鸟类
type bird1 struct {
}

// 实现飞行动物接口
func (b *bird1) Fly() {
	fmt.Println("bird: fly")
}

func main() {
	/**
	空接口类型
	1、可以使用interface{}表示任何类型，空接口类型可以保存任何值，也可以从空接口中取出原值。
	*/
	var any interface{} = 1
	fmt.Println(any) //1

	any = "hello"
	fmt.Println(any) //hello

	//var i interface{} = 1
	// 声明b变量, 尝试赋值i
	//var b int = i//报错，`interface{}`较大，`int`较小，不可向下兼容

	/**
	类型断言
	1、类型断言是一个使用在接口值上的操作，用于检查接口类型变量所持有的值是否实现了期望的接口或者具体的类型
	2、value, ok := x.(T) 其中，x 表示一个接口的类型，T 表示一个具体的类型（也可为接口类型）。
	*/
	var x interface{}
	x = 10
	value, ok := x.(int)
	fmt.Print(value, ",", ok) //输出10,true

	value2, ok := x.(string)
	fmt.Print(value2, ",", ok) //输出 ,false

	//value3 := x.(string)
	//fmt.Println(value3)//报错，interface {} is int, not string

	//判断接口类型是否符合
	var aa interface{} = &bird1{}
	value3, ok3 := aa.(Flyer1)
	fmt.Println(value3, ok3) //&{} true

	getType(10) //the type of a is int
}

func getType(a interface{}) {
	//接口变量.(type) 表示会判断输出值的类型，比如int、string、float64
	switch a.(type) {
	case int:
		fmt.Println("the type of a is int")
	case string:
		fmt.Println("the type of a is string")
	case float64:
		fmt.Println("the type of a is float")
	default:
		fmt.Println("unknown type")
	}
}
