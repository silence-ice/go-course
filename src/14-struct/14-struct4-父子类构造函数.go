package main

import "fmt"

type CatA struct {
	Color string
	Name  string
}

type BlackCat struct {
	//嵌入了 Cat 结构体，BlackCat 拥有 Cat 的所有成员，实例化后可以自由访问 Cat 的所有成员
	CatA
}

func main() {
	//fmt.Println(NewCat("huangbin"))
	fmt.Println(NewBlackCat("black"))
}

// “构造子类”
func NewBlackCat(color string) *BlackCat {
	cat := &BlackCat{}
	cat.Color = color
	return cat
}
