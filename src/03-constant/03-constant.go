package main

import (
	"fmt"
	"unsafe"
)

/**
 * 1、常量是不会被修改的量，数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型
 * 2、常量一般和普通变量声明即可，不用大写，首字母大写表示public
 */
const name string = "huangbin" //显式声明
const age = 13                 //隐式声明
const (
	mama   string = "mama"
	father string = "father"
)
const apple, banana string = "苹果", "香蕉"

func main() {
	fmt.Println(name, age)
	fmt.Println(len(name))                   //8
	fmt.Println(unsafe.Sizeof(name))         //16
	fmt.Println(mama, father, apple, banana) //mama father 苹果 香蕉

	//常量枚举
	const (
		Unknown = 0
		Female  = 1
		Male    = 2
	)
	println(Unknown, Female, Male)

	const (
		o = iota
		p = iota
		q = iota
	)
	fmt.Println(o, p, q) //0 1 2

	const (
		m = iota
		w = iota
		_
		_
		x = iota
	)
	fmt.Println(m, w, x) //0 1 4

	//iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次
	//iota 可理解为 const 语句块中的行索引
	//若常量没有声明值，就会使用iota值
	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i) //0 1 2 ha ha 100 100 7 8
}
