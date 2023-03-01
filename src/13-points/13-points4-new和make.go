package main

import "fmt"

/*
*
make和new的区别
1、make 关键字的主要作用是创建 slice、map 和 channel 等内置的数据结构
2、new 的主要作用是为类型申请一片内存空间，并返回指向这片内存的指针
*/
func main() {
	type Student struct {
		name string
		age  int
	}

	var s *Student
	s = new(Student) //分配内存空间，相当于 &Student{}
	fmt.Println(s)   //&{ 0}
	s.name = "dequan"
	fmt.Println(s.name) //dequan

	//make 也是用于内存分配的，但是和 new 不同，它只用于 chan、map 以及 slice 的内存创建，而且它返回的类型就是这三个类型本身
	countryCapitalMap := make(map[string]string) //创建集合
	s2 := make([]int, 16)                        //[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0], len=16, cap=16
	fmt.Println(countryCapitalMap)               //map[]
	fmt.Println(s2)                              //[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]

	/**
	Go语言中的 new 和 make 主要区别如下：
	1、make 只能用来分配及初始化类型为 slice、map、chan 的数据。new 可以分配任意类型的数据；
	2、new 分配返回的是指针，即类型 *Type。make 返回引用，即 Type；
	3、new 分配的空间被清零。make 分配空间后，会进行初始化；
	*/
}
