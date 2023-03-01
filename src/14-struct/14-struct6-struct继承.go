package main

import "fmt"

/**
结构体内嵌 模拟 类的继承
1、人类不能飞行，鸟类可以飞行。人类和鸟类都可以继承自可行走类，但只有鸟类继承自飞行类。
2、Go语言的结构体内嵌特性就是一种组合特性，使用组合特性可以快速构建对象的不同特性。
*/
// 可飞行的
type Flying struct{}

func (f *Flying) Fly() {
	fmt.Println("can fly")
}

// 可行走的
type Walkable struct{}

func (f *Walkable) Walk() {
	fmt.Println("can calk")
}

type Human struct {
	Walkable // 人类能行走
}

type Bird struct {
	Walkable // 鸟类能行走
	Flying   // 鸟类能飞行
}

func main() {
	//实例化鸟类
	b := new(Bird)
	b.Fly()  //can fly
	b.Walk() //can calk

	//实例化人类
	h := new(Human)
	h.Walk() //can calk
}
