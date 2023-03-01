package main

import (
	"fmt"
)

func main() {
	//one 数据类型别名
	type Huangbin int32
	var h Huangbin = 123
	fmt.Println(h) //123

	//two 字典定义
	type person2 map[string]string
	person2Map := person2{
		"name": "huangbin",
		"age":  "23",
	}

	//功能同上
	//person := map[string]string {
	//	"name": "huangbin",
	//	"age": "23",
	//}

	for one := range person2Map {
		fmt.Println(person2Map[one])
	}

	//three 结构体定义
	type person struct {
		name string //注意后面不能有逗号
		age  int
	}

	//结构体初始化
	p := person{
		name: "taozs", //注意后面要加逗号
		age:  18,      //或者下面的}提到这儿来可以省略逗号
	}
	fmt.Println(p.name) //taozs

	//four 接口定义
	type Phone interface {
		call(str string)
	}

	//four 函数定义
	type handle func(str string) //自定义一个函数func，别名叫做handle，传入一个string参数
}
