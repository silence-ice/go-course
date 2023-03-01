package main

import "fmt"

func main() {
	//1、条件判断
	var age int = 23
	if age == 25 {
		fmt.Println("true")
	} else if age < 25 {
		fmt.Println("too small")
	} else {
		fmt.Println("too big")
	}

	//2、￿switch会自动break，除非使用fallthrough
	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}
	fmt.Println(grade) //A

	//3、switch一分支多值
	var a = "mum"
	switch a {
	case "mum", "daddy":
		fmt.Println("family") //family
	}

	//4、switch——分支表达式
	var r int = 11
	switch {
	case r > 10 && r < 20:
		fmt.Println(r) //11
	}

	//5、￿switch——使用fallthrough
	var s = "hello"
	switch {
	case s == "hello":
		fmt.Println("hello") //输出hello
		fallthrough
	case s != "world":
		fmt.Println("world") //输出world
	}

}
