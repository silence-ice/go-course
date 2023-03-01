package main

import (
	"fmt"
	"strconv"
)

/*
*
strconv包实现了基本数据类型与其字符串表示的转换，主要有以下常用函数： Atoi()、Itia()、parse系列、format系列、append系列。
*/
func main() {
	//Atoi()函数用于将字符串类型的整数转换为int类型
	s1 := "100"
	i1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("can't convert to int")
	} else {
		fmt.Printf("type:%T value:%#v\n", i1, i1) //type:int value:100
	}

	//Itoa()函数用于将int类型数据转换为对应的字符串表示
	i2 := 200
	s2 := strconv.Itoa(i2)
	fmt.Printf("type:%T value:%#v\n", s2, s2) //type:string value:"200"

	/**
	Parse类函数用于转换字符串为给定类型的值：ParseBool()、ParseFloat()、ParseInt()、ParseUint()

	1、func ParseBool(str string) (value bool, err error)

	2、func ParseInt(s string, base int, bitSize int) (i int64, err error)
		2.1 返回字符串表示的整数值，接受正负号。
		2.2 base指定进制（2到36），如果base为0，则会从字符串前置判断，”0x”是16进制，”0”是8进制，否则是10进制；
		2.3 bitSize指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64；

	3、func ParseUint(s string, base int, bitSize int) (n uint64, err error)
		3.1 ParseUint类似ParseInt但不接受正负号，用于无符号整型。

	4、func ParseFloat(s string, bitSize int) (f float64, err error)
		4.1 解析一个表示浮点数的字符串并返回其值。
		4.2 bitSize指定了期望的接收类型，32是float32（返回值可以不改变精确值的赋值给float32），64是float64；
	*/
	b, err := strconv.ParseBool("true")
	f, err := strconv.ParseFloat("3.1415", 64)
	i, err := strconv.ParseInt("-2", 10, 64)
	u, err := strconv.ParseUint("2", 10, 64)
	fmt.Printf("type:%T value:%#v\n", b, b) //type:bool value:true
	fmt.Printf("type:%T value:%#v\n", f, f) //type:float64 value:3.1415
	fmt.Printf("type:%T value:%#v\n", i, i) //type:int64 value:-2
	fmt.Printf("type:%T value:%#v\n", u, u) //type:uint64 value:0x2

	/**
	Format系列函数实现了将给定类型数据格式化为string类型数据的功能

	1、func FormatBool(b bool) string

	2、func FormatInt(i int64, base int) string
		2.1 返回i的base进制的字符串表示。base 必须在2到36之间，结果中会使用小写字母’a’到’z’表示大于10的数字。

	3、func FormatUint(i uint64, base int) string
		3.1 是FormatInt的无符号整数版本。

	4、func FormatFloat(f float64, fmt byte, prec, bitSize int) string
		4.1 函数将浮点数表示为字符串并返回。
		4.2 bitSize表示f的来源类型（32：float32、64：float64），会据此进行舍入。
		4.3 fmt表示格式：’f’（-ddd.dddd）、’b’（-ddddp±ddd，指数为二进制）、’e’（-d.dddde±dd，十进制指数）、’E’（-d.ddddE±dd，十进制指数）、’g’（指数很大时用’e’格式，否则’f’格式）、’G’（指数很大时用’E’格式，否则’f’格式）。
		4.4 prec控制精度（排除指数部分）：对’f’、’e’、’E’，它表示小数点后的数字个数；对’g’、’G’，它控制总的数字个数。如果prec 为-1，则代表使用最少数量的、但又必需的数字来表示f。
	*/
	s3 := strconv.FormatBool(true)
	s4 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	s5 := strconv.FormatInt(-2, 16)
	s6 := strconv.FormatUint(2, 16)
	fmt.Printf("type:%T value:%#v\n", s3, s3) //type:string value:"true"
	fmt.Printf("type:%T value:%#v\n", s4, s4) //type:string value:"3.1415E+00"
	fmt.Printf("type:%T value:%#v\n", s5, s5) //type:string value:"-2"
	fmt.Printf("type:%T value:%#v\n", s6, s6) //type:string value:"2"
}
