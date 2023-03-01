package main

import "fmt"

func main() {
	/**
	1、数组是一个由固定长度的特定类型元素组成的序列，一个数组可以由零个或多个元素组成。因为数组的长度是固定的，所以在Go语言中很少直接使用数组。
	2、和数组对应的类型是 Slice（切片），Slice 是可以增长和收缩的动态序列，功能也更灵活，我们先暂时学习数组。
	3、数组的声明方式：var variable_name [SIZE] variable_type
	*/
	//声明方法
	var a [10]int
	fmt.Println(a) //[0 0 0 0 0 0 0 0 0 0]

	//简洁声明方法，但是需要有默认值
	z := [2]int{1, 2}
	fmt.Println(z) //[1 2]

	//初始化数组（限定个数）
	var b = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println(b)    //[1000 2 3.4 7 50]
	fmt.Println(b[1]) //2

	//初始化数组（不限定个数）
	var c = [...]float32{1000.0, 2.0, 3.4, 7.0}
	fmt.Println(c) //[1000 2 3.4 7 50]
	fmt.Println(c[:])
	fmt.Println(c[0:2]) //[1000 2]
	fmt.Println(c[2:])  //[3.4 7 50]

	////////////////////////////////////////////////////////////////////////////////////
	//遍历方式一
	for i := 0; i < len(c); i++ {
		fmt.Println(c[i]) //1000.0, 2.0, 3.4, 7.0, 50.0
	}

	//遍历方式二
	for j := range c {
		fmt.Println(c[j]) //1000.0, 2.0, 3.4, 7.0, 50.0
	}

	//遍历方式三
	for k, v := range c {
		fmt.Println(k, v) //输出key value
	}

	//遍历方式四
	for _, v := range c {
		fmt.Println(v) //只要value，不要key
	}

	//遍历方式五-函数封装
	printArray(b) //正常输出
	//printArray(z)//错误输出，因为函数只要float32类型的
	//printArray(c)//错误输出，必须满足[5]float32的条件，现在只有[4]float32

	//遍历方式六-函数的传参引用
	printRightArray(&b)

	//遍历方式七-功能同 遍历方式六 一样
	printSileceArray(b[:]) //！！！切片默认是引用传递！！！

	////////////////////////////////////////////////////////////////////////////////////

	//初始化二维数组
	var d = [3][4]int{
		{0, 1, 2, 3},   /*  第一行索引为 0 */
		{4, 5, 6, 7},   /*  第二行索引为 1 */
		{8, 9, 10, 11}, /* 第三行索引为 2 */
	}
	fmt.Println(d)       //[1000 2 3.4 7 50]
	fmt.Println(d[1][2]) //6

	//// 声明一个 2×2 的二维整型数组
	var array [2][2]int
	// 设置每个元素的整型值
	array[0][0] = 10
	array[0][1] = 20
	array[1][0] = 30
	array[1][1] = 40
	fmt.Println(array) //[[10 20] [30 40]]

}

func printArray(arr [5]float32) {
	arr[0] = 12313.1 //不会改变外面的输出，因为在go中是值传递（复制arr数组），而非引用传递
	for k, v := range arr {
		fmt.Println(k, v) //输出key value
	}
}

/**
 * 指针函数
 */
func printRightArray(arr *[5]float32) {
	arr[0] = 12313.1 //可以改变外面的输出
	for k, v := range arr {
		fmt.Println(k, v) //输出key value
	}
}

func printSileceArray(arr []float32) {
	arr[0] = 12313.1 //可以改变外面的输出
	for k, v := range arr[:] {
		fmt.Println(k, v) //输出key value
	}
}
