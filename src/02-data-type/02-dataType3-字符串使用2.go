package main

import (
	"fmt"
	"strings"
)

func main() {
	//one 字符串截取
	//截取下面字符串`死神来了, 死神bye bye`中的`死神bye bye`
	tracer := "死神来了, 死神bye bye"
	comma := strings.Index(tracer, ",")
	fmt.Println(comma) //输出12，一个中文字符算三个字节

	//分开截取的目的，主要是有两个 死神的字符串
	tempStr := tracer[comma:]
	fmt.Println(tempStr) //输出 `, 死神bye bye`
	pos := strings.Index(tempStr, "死神")
	fmt.Println(pos) //输出2

	fmt.Println(tracer[comma+pos:]) //输出`死神bye bye`

	//two 字符串遍历
	theme := "狙击 a"
	for i := 0; i < len(theme); i++ {
		/**
		ascii: ç  231
		ascii: å  139
		ascii: å  153
		ascii: å  229
		ascii: å  135
		ascii: »  187
		ascii:    32
		ascii: a  97
		*/
		fmt.Printf("ascii: %c  %d\n", theme[i], theme[i])
	}

	//结论：上面的遍历方式只适合全ASCII字符的遍历，如果有中文还是需要按Unicode字符遍历字符串

	for _, val := range theme {
		/**
		Unicode: 狙  29401
		Unicode: 击  20987
		Unicode:    32
		Unicode: a  97
		*/
		fmt.Printf("Unicode: %c  %d\n", val, val)
	}

	//修改字符串
	//Go 语言的字符串无法直接修改每一个字符元素，只能通过重新构造新的字符串并赋值给原来的字符串变量实现
	angel := "Heros never die"
	angleBytes := []byte(angel)

	for i := 6; i <= 7; i++ {
		angleBytes[i] = ' '
	}
	fmt.Println(string(angleBytes))

	/**
	总结
		1、Go 语言的字符串是不可变的。
		2、修改字符串时，可以将字符串转换为 []byte 进行修改。
		3、[]byte 和 string 可以通过强制类型转换互转
	*/

}
