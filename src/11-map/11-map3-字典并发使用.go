package main

import (
	"fmt"
	"sync"
)

/*
*
sync.Map的出现
1、Go语言中的 map 在并发情况下，只读是线程安全的，同时读写是线程不安全的。
2、需要并发读写时，一般的做法是加锁，但这样性能并不高，Go语言在 1.9 版本中提供了一种效率较高的并发安全的 sync.Map

sync.Map 有以下特性：
1、无须初始化，直接声明即可。
2、sync.Map 不能使用 map 的方式进行取值和设置等操作，而是使用 sync.Map 的方法进行调用，Store 表示存储，Load 表示获取，Delete 表示删除。
3、使用 Range 配合一个回调函数进行遍历操作，通过回调函数返回内部遍历出来的值，Range 参数中回调函数的返回值在需要继续迭代遍历时，返回 true，终止迭代遍历时，返回 false。
*/
func main() {
	var scene sync.Map

	// 将键值对保存到sync.Map
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)

	// 从sync.Map中根据键取值
	resA, resB := scene.Load("london")
	fmt.Println(resA) //100
	fmt.Println(resB) //true

	// 根据键删除对应的键值对
	scene.Delete("london")

	//遍历所有sync.Map中的键值对
	//Range() 方法可以遍历 sync.Map，遍历需要提供一个匿名函数，参数为 k、v，类型为 interface{}
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})

}
