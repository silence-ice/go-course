package main

import "fmt"

/*
*
给struct添加方法的格式如下：

	func (接收器变量 接收器类型) 方法名(参数列表) (返回参数) {
		函数体
	}

	1、接收器变量：接收器中的参数变量名在命名时，官方建议使用接收器类型名的第一个小写字母，而不是 self、this 之类的命名
	2、接收器类型：接收器类型和参数类似，可以是指针类型和非指针类型。
	3、方法名、参数列表、返回参数：格式与函数定义一致。
*/
type Property struct {
	value int // 属性值
}

/**
1、这个方法是实现了fmt包中print.go接口
	type Stringer struct{
		String() string
	}
2、这个方法常用于触发了fmt.Print方法后，会将 输出 struct 的值 变成成 该方法返回的的数据
*/
//func (p *Property) String() string {
//	return "huangbin"
//}

// 设置属性值
// 1、指针类型的接收器由一个结构体的指针组成，支持修改接收器指针的任意成员变量。
// 2、非指针接收器的方法中可以获取接收器的成员值，但不允许修改。
func (p *Property) SetValue(v int) {
	p.value = v // 修改p的成员变量
}

// 取属性值
func (p *Property) Value() int {
	return p.value
}

// 虽然非指针类型的接收器不能修改 Property 的成员属性，但是可以在计算后返回新的 Property 对象
// 返回类型是 Property 的struct
func (p Property) NewProperty() Property {
	return Property{
		value: 99,
	}
}

func main() {
	//指针型方法测试
	p := new(Property)
	//p := &Property{}
	fmt.Println(p) //&{0}

	p.SetValue(99)
	fmt.Println(p) //&{99}，返回引用地址

	//非指针型方法测试
	pp := new(Property)
	fmt.Println(pp.NewProperty()) //{99}，返回Property值
}
