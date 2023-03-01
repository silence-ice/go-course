package main

import (
	"fmt"
	"strconv"
)

func main() {
	//for 循环
	sum := 0

	for i := 1; i <= 5; i++ {
		sum += i
	}
	fmt.Println(sum) //输出15

	//for 循环——忽略初始语句，在for表达式外进行声明
	k := 1
	for ; k <= 3; k++ {
		sum += k
	}
	fmt.Println(sum) //输出21=15+6

	//for 循环——忽略条件表达式，在for表达式里面声明
	p := 1
	for ; ; p++ {
		if p >= 3 {
			break
		}
		sum += p
	}
	fmt.Println(sum) //输出24=21+3

	//功能同上，将累加表达式放在for表达式里面声明
	z := 1
	for {
		if z >= 3 {
			break
		}
		sum = sum + z
		z++
	}
	fmt.Println(sum) //输出27=24+3

	//2、for 循环
	result := ""
	//n => 5 => 2.5 => 1.25 => 0.625
	for n := 5; n > 0; n = n / 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	fmt.Println(result) //输出101

	//3、while循环
	sum2 := 1
	for sum2 <= 5 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	//4、数组循环
	var strings = []string{"google", "runoob"}
	//strings := []string{"google", "runoob"}
	for i, s := range strings {
		fmt.Println(i, s)
	}

	numbers := [6]int{1, 2, 3, 5}
	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}

	//5、循环
	var a int = 10
LOOP:
	for a < 17 {
		if a == 15 {
			/* 跳过迭代 */
			a = a + 1
			goto LOOP
		}
		fmt.Printf("a的值为 : %d\n", a) //输出10～16，15是直接跳过的
		a++
	}
}
