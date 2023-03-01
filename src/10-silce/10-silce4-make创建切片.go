package main

import "fmt"

func main() {
	fmt.Println("Creating slice")
	var s []int // Zero value for slice is nil

	for i := 0; i < 10; i++ {
		//下面证明容量是呈2倍进行增长的
		//[], len=0, cap=0
		//[1], len=1, cap=1
		//[1 3], len=2, cap=2
		//[1 3 5], len=3, cap=4
		//[1 3 5 7], len=4, cap=4
		//[1 3 5 7 9], len=5, cap=8
		//[1 3 5 7 9 11], len=6, cap=8
		//[1 3 5 7 9 11 13], len=7, cap=8
		//[1 3 5 7 9 11 13 15], len=8, cap=8
		//[1 3 5 7 9 11 13 15 17], len=9, cap=16
		printSlice(s) //容量会以2倍基数增加
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	printSlice(s1) //[2 4 6 8], len=4, cap=4

	//方式二——使用make创建切片
	//make函数格式如下：make( []Type, size, cap )
	s2 := make([]int, 16)     //[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0], len=16, cap=16
	s3 := make([]int, 10, 32) //[0 0 0 0 0 0 0 0 0 0], len=10, cap=32
	printSlice(s2)
	printSlice(s3)

	fmt.Println("Copying slice")
	copy(s2, s1) //[2 4 6 8 0 0 0 0 0 0 0 0 0 0 0 0], len=16, cap=16
	printSlice(s2)

	fmt.Println("Deleting elements from slice")
	s2 = append(s2[:3], s2[4:]...) //把第三个元素删除掉，把开始到第个2的元素 和 第4个元素到末尾 用 ...（语法）进行链接
	printSlice(s2)

	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]

	fmt.Println(front)
	printSlice(s2)

	fmt.Println("Popping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]

	fmt.Println(tail)
	printSlice(s2)
}

func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n", s, len(s), cap(s))
}
