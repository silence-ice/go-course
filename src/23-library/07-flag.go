package main

import (
	"flag"
	"fmt"
	"time"
)

/**
 * Go语言内置的 flag 包实现了命令行参数的解析，flag 包使得开发命令行工具更为简单。
 */
func main() {
	//mode1()
	mode2()
}

/**
 * flag.Type(flag 名, 默认值, 帮助信息) *Type
 * 1、Type 可以是 Int、String、Bool 等，返回值为一个相应类型的指针
 * 以下面举例：go run 07-flag.go -name "silence"
 * 2、第一个参数是flag名，对标name
 * 3、第二个参数是默认值，对标silence
 * 2、第三个参数是帮助信息，用于备注
 */
func mode1() {
	//go run 07-flag.go -name "silence" -age 27

	name := flag.String("name", "张三", "姓名")
	age := flag.Int("age", 18, "年龄")
	married := flag.Bool("married", false, "婚否")
	delay := flag.Duration("d", 0, "时间间隔")

	//通过以上两种方法定义好命令行 flag 参数后，需要通过调用 flag.Parse() 来对命令行参数进行解析
	flag.Parse()
	fmt.Println(*name, *age, *married, *delay) //silence 27 false 0s
}

/**
 * flag.TypeVar(Type 指针, flag 名, 默认值, 帮助信息)
 * 1、和flag.Type()不同的是第一个参数是变量，而flag.TypeVar()第一个参数是指针
 */
func mode2() {
	//go run 07-flag.go -name "silence" -age 27

	var name string
	var age int
	var married bool
	var delay time.Duration
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "时间间隔")

	flag.Parse()

	fmt.Println(name, age, married, delay) //silence 27 false 0s
}
