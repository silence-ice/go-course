package main

import "fmt"

func main() {
	//方式一——使用数组创建切片
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7}

	s1 := arr[2:6]
	s2 := s1[3:5]

	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1)) //s1=[2 3 4 5], len(s1)=4, cap(s1)=6
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2)) //s2=[5 6], len(s2)=2, cap(s2)=3
	fmt.Println(arr)                                                    //[0 1 2 3 4 5 6 7]

	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s3, len(s3), cap(s3)) //s2=[5 6 10], len(s2)=3, cap(s2)=3
	//添加元素时如果超越cap，系统会重新分配更大的底层数组,以2倍基数增长容量
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s4, len(s4), cap(s4)) //s2=[5 6 10 11], len(s2)=4, cap(s2)=6
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s5, len(s5), cap(s5)) //s2=[5 6 10 11 12], len(s2)=5, cap(s2)=6

	fmt.Println(arr) //[0 1 2 3 4 5 6 10]

}
