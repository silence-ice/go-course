package main

import "fmt"

func main() {
	//遍历元素
	slice := []int{10, 20, 30, 40}
	// range 返回的是每个元素的副本，而不是直接返回对该元素的引用
	for index, value := range slice {
		fmt.Printf("Index: %d Value: %d\n", index, value)
	}

	//添加元素
	var a []int

	a = append(a, 1)                 // 追加1个元素
	a = append(a, 1, 2, 3)           // 追加多个元素, 手写解包方式
	a = append(a, []int{1, 2, 3}...) // 追加一个切片, 切片需要解包
	fmt.Println(a)                   //[1 1 2 3 1 2 3]

	//复制元素
	//copy() 函数的第一个参数是要复制的目标 slice，第二个参数是源 slice，两个 slice 可以共享同一个底层数组，甚至有重叠也没有问题
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	fmt.Println(slice2)  //输出[1 2 3]
	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
	fmt.Println(slice1)  //输出[1 2 3 4 5] 因为silce2被重置了，所以silce1是不会变得

	//删除元素
	//1、从开头删除元素
	b := []int{1, 2, 3}
	b = b[1:]      // 删除开头1个元素
	fmt.Println(b) //[2 3]

	//2、从中间位置删除
	c := []int{1, 2, 3, 4}
	//append方法第一个参数是获取0:2索引之间的值，第二个参数是获取3到最后的值，最后进行合并
	c = append(c[:2], c[2+1:]...) //输出[1 2 4]，将第二个索引进行删除
	fmt.Println(c)

	//3、从尾部位置删除
	d := []int{1, 2, 3, 4}
	d = d[:len(d)-1] //输出[1 2 3]，删除尾部1个元素
	d = d[:len(d)-2] //输出[1]，删除尾部2个元素
	fmt.Println(d)

}
