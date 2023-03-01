package main

import "fmt"

func main() {
	//nil 是 map、slice、pointer、channel、func、interface 的零值
	var m map[int]string
	var ptr *int
	var c chan int
	var sl []int
	var f func()
	var i interface{}
	fmt.Printf("%#v\n", m)
	fmt.Printf("%#v\n", ptr)
	fmt.Printf("%#v\n", c)
	fmt.Printf("%#v\n", sl)
	fmt.Printf("%#v\n", f)
	fmt.Printf("%#v\n", i)

	/**
	map[int]string(nil)
	(*int)(nil)
	(chan int)(nil)
	[]int(nil)
	(func())(nil)
	<nil>
	*/
}
