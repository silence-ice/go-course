package main

import (
	"errors"
	"fmt"
)

func main() {
	/**
	panic: crash
	goroutine 1 [running]:
	main.main()
	        /Users/silence/Documents/go/course/src/17-error/18-painc.go:4 +0x39
	exit status 2
	*/
	//panic("crash")

	/**
	宕机后要做的事情2
	宕机后要做的事情1
	panic: 宕机

	goroutine 1 [running]:
	main.main()
	        /Users/silence/Documents/go/course/src/17-error/18-painc.go:18 +0xf1
	exit status 2
	*/
	defer fmt.Println("宕机后要做的事情1")
	defer fmt.Println(errors.New("宕机后要做的事情2"))
	panic("宕机")
}
