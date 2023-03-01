package main

import "fmt"

func main() {
	//指针数组
	const MAX int = 3
	b := []int{10, 100, 200}
	var i int
	var ptr2 [MAX]*int

	for i = 0; i < MAX; i++ {
		ptr2[i] = &b[i] /* 整数地址赋值给指针数组 */
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr2[i]) //10 100 200
	}

	//如果一个指针变量存放的又是另一个指针变量的地址，则称这个指针变量为指向指针的指针变量,声明格式如下：var ptr **int;
	//ptr2 -> ptr ->变量a
	var a int
	var ptr *int
	var pptr **int

	a = 3000

	//指针 ptr 地址
	ptr = &a

	//指向指针 ptr 地址
	pptr = &ptr

	/* 获取 pptr 的值 */
	fmt.Printf("变量 a = %d\n", a)                      //3000
	fmt.Printf("指针变量 *ptr = %d\n", *ptr)              //3000
	fmt.Printf("指向指针的指针内存地址 **pptr = %d\n", *pptr)    //824634236952
	fmt.Printf("指向指针的指针变量对应的值 **pptr = %d\n", **pptr) //3000

}
