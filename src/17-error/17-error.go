package main

import (
	"errors"
	"fmt"
	"math"
)

// Go语言中引入 error 接口类型作为错误处理的标准模式，如果函数要返回错误，则返回值类型列表中肯定包含 error。
// 创建一个 error 最简单的方法就是调用 errors.New 函数，它会根据传入的错误信息返回一个新的 error。
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return -1, errors.New("math: square root of negative number")
	}
	return math.Sqrt(f), nil
}

func main() {
	//result, err := Sqrt(13)//3.605551275463989
	result, err := Sqrt(-13) //math: square root of negative number

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
