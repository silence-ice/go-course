package main

import "fmt"

/*
*
1、Go 语言切片是对数组的抽象，之后在开发中只需要使用切片即可
2、Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，功能强悍的内置类型切片("动态数组")
3、切片与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

切片的结构
1、Go语言中切片的内部结构包含地址、大小和容量，切片一般用于快速地操作一块数据集合。
2、如果将数据集合比作切糕的话，切片就是你要的“那一块”，切的过程包含从哪里开始（切片的起始位置）及切多大（切片的大小），容量可以理解为装切片的口袋大小
*/
func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	//points 获取指定范围的值
	fmt.Println("arr[2:6] =", arr[2:6]) //arr[2:6] = [2 3 4 5]
	fmt.Println("arr[:6] =", arr[:6])   //arr[:6] = [0 1 2 3 4 5]
	fmt.Println("arr[:6] =", arr[0:0])  //重置切片，清空拥有的元素

	////////////////////////////////////////////////////////

	//points 切片赋值、更改索引值测试
	s1 := arr[2:]
	fmt.Println("s1 =", s1) //s1 = [2 3 4 5 6 7]
	s2 := arr[:]            //复制原有的切片
	fmt.Println("s2 =", s2) //s2 = [0 1 2 3 4 5 6 7]

	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)
	//one 证明切片可以修改值
	fmt.Println(s1) //[100 3 4 5 6 7]
	//two 原变量也被修改了，但是修改的值是在index=2，因为s1赋值的时候是从index=2进行赋值的
	fmt.Println(arr) //[0 1 100 3 4 5 6 7]

	////////////////////////////////////////////////////////

	//points 切片赋值、更改索引值测试
	fmt.Println("After updateSlice(s2)")
	//three s2也发生了改变，因为原变量都变了，其他引用的变量都会跟着变
	fmt.Println(s2) //[0 1 100 3 4 5 6 7]
	updateSlice(s2)
	fmt.Println(s2)  //[100 1 100 3 4 5 6 7]
	fmt.Println(arr) //[100 1 100 3 4 5 6 7]

	////////////////////////////////////////////////////////

	fmt.Println("Reslice")
	fmt.Println(s2) //[100 1 100 3 4 5 6 7]
	s2 = s2[:5]
	fmt.Println(s2) //[100 1 100 3 4]
	s2 = s2[2:]
	fmt.Println(s2) //[100 3 4]

	fmt.Println("Extending slice")
	arr[0], arr[2] = 0, 2
	fmt.Println("arr =", arr) //arr = [0 1 2 3 4 5 6 7]
	s1 = arr[2:6]
	s2 = s1[3:5]

	/**
	 * silce可以向后扩展，不可以向前扩展
	 * s[i]不可以超越len(s)，向后扩展不可以超越底层数组cap(s)
	 */
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1)) //s1=[2 3 4 5], len(s1)=4, cap(s1)=6；其中4表示长度，6表示容量；且规则容量必须要 > 长度

	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2)) //s2=[5 6], len(s2)=2, cap(s2)=3

	////////////////////////////////////////////////////////

	// 声明一个二维整型切片并赋值
	slice := [][]int{{10}, {100, 200}}
	slice[0] = append(slice[0], 20)
	fmt.Println(slice) //[[10 20] [100 200]]
}

func updateSlice(s []int) {
	s[0] = 100
}
