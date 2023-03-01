package main

import "fmt"

var (
	oneQuestion3 string
)

func main() {
	//one
	fmt.Printf("请输入： ")
	fmt.Scanln(&oneQuestion3) //Scanln 扫描来自标准输入的文本，将空格分隔的值依次存放到后续的参数内，直到碰到换行。

}
