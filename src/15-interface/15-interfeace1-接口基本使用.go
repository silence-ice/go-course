package main

import "fmt"

/**
什么是Duck Typing?
	1、Duck Typing 的原话是，走起来像鸭子、游泳起来像鸭子、叫起来也像鸭子，那么它就是一只鸭子。
	2、这个原话是可以灵活理解的，就看我们怎么定义鸭子的行为，我们可以说，能浮在水上游的，黄色的，可爱的就是鸭子，那么，图片中的大黄鸭，它就是一只鸭子！
	3、这就是所谓的 Duck Typing，它只关心事物的外部行为而非内部结构。它并不关心你这只鸭子是长肉的还是充气的。在编程中，也常常用这种方式来描述事物。
*/

/*
*
接口声明的格式

	1、每个接口类型由数个方法组成。接口的形式代码如下：
		type 接口类型名 interface{
			方法名1( 参数列表1 ) 返回值列表1
			方法名2( 参数列表2 ) 返回值列表2
			…
		}
*/
type Receiver interface {
	Get(url string) string //声明Get方法中url变量类型为string，返回类型也是string
}

////////////////////////////////实现方//////////////////////////////////

// 定义 MyReceiver 结构
// 可以看到下面实现方是没有出现任何接口的信息
type MyReceiver struct {
	Content string
}

// 要使用 Receiver 接口，就必须实现接口的Get方法，方法中的名称、参数列表、返回参数列表都必须一样
// func (r *MyReceiver) Get(url string) string{//指针引用方式
func (r MyReceiver) Get(url string) string {
	return r.Content + "  " + url
}

func main() {
	//！！！！接口一定要用结构体实现！！！
	//r是一个接口变量，表示声明 Receiver 接口的 r 变量
	var r Receiver

	//将实现好的 MyReceiver 结构赋值给接口变量
	//虽然 r 是一个接口变量，且 MyReceiver 已经完全实现了 Receiver 的所有方法，因此赋值是成功的。
	//r = &MyReceiver{Content:"huangbin"}//指针引用方式
	r = MyReceiver{Content: "huangbin"} //

	//虽然功能同上，但是没有经过接口变量的声明，其实就是自己实现func自己调用而已
	//r := MyReceiver{Content:"huangbin"}

	//调用方式一
	fmt.Println(r.Get("www.baidu.com")) //huangbin  www.baidu.com

	//调用方式二 封装方法
	fmt.Println(download(r)) //huangbin  www.imooc.com
}

////////////////////////////////封装方法//////////////////////////////////
/**
声明 download 方法，url变量类型是一个 Receiver 对象，返回类型是string
*/
func download(url Receiver) string {
	return url.Get("www.imooc.com") //使用 Receiver 对象中的 Get 方法
}
