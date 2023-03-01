package main

import "fmt"

/*
*
func有三种声明方式：
1、func add(x int, y int) (int) {return x + y} => 正常声明方式
2、func add(x int, y int) int {return x + y} => 功能同上
3、func sub(x, y int) (z int) { z = x - y; return} => 返回类型带上变量(z)，函数体内return就无须带任何数据，默认返回变量z
*/
func main() {
	/* 定义局部变量 */
	var a int = 300
	var b int = 200
	var ret int

	/* 调用函数并返回最大值 */
	ret = max(a, b)
	fmt.Printf("最大值是 : %d\n", ret)

	//返回多个值
	fruitA1, fruitB1 := listFruit("apple", "banana")
	println(fruitA1, fruitB1) //apple banana

	//返回错误信息
	//result,err := eval(3, 4, "xxxx");
	result, err := eval(3, 4, "*")
	if err != nil {
		fmt.Println(err) //unsupported operation: x
	} else {
		fmt.Println(result) //12
	}
}

/*
*

	函数返回两个数的最大值
	func function_name( [parameter list] ) [return_types] {函数体}
*/
func max(num1 int, num2 int) int {
	/* 定义局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

/**
 * 函数返回多个值
 */
func listFruit(param1 string, param2 string) (string, string) {
	return param1, param2 //返回多个值
}

/**
 * 函数返回错误信息
 */
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b) //如果不要返回值，可以用_来取代进行返回
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

func div(a, b int) (int, int) {
	return a / b, a % b
}
