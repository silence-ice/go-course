package main

import "fmt"

type treeNode struct {
	value int
	Left  *treeNode //声明指针式变量
	Right *treeNode
}

/**
 * [功能同下]给treeNode struct声明一个方法，打印value
 */
func (node treeNode) Print() {
	fmt.Println(node.value)
}

/**
 * [功能同上]声明一个方法，打印treeNode struct的value
 */
func printValue(node treeNode) {
	fmt.Print(node.value)
}

/**
 * 因为函数的传参是传值的，所以在函数里面修改值是没有用的
 */
func (node treeNode) SetValue(value int) {
	node.value = value
}

/**
 * 修改为指针传递，有效修改值
 * 1、一般要改变内容就需要使用指针接受者
 * 2、结构过大也考虑使用指针接收者
 */
func (node *treeNode) SetTrueValue(value int) {
	node.value = value
}

/**
 * 给treeNode struct声明中序遍历方法，中序遍历顺序是 左->中->右
 */
func (node *treeNode) Traverse() {
	if node == nil {
		return
	}

	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

func main() {
	var root treeNode //引用treeNode结构体

	root.value = 3
	root.Left = &treeNode{} //因为left被声明为一个指针，所需要需要用引用进行声明
	root.Right = &treeNode{value: 5, Left: nil, Right: nil}

	//new(T) 为每个新的类型T分配一片内存，初始化为 0 并且返回类型为*T的内存地址；它相当于 &T{}，它适用于值类型如数组和结构体
	//make(T) 返回一个类型为 T 的初始值，它只适用于3种内建的引用类型：切片、map 和 channel
	root.Right.Left = new(treeNode) //功能同`root.right.left = &treeNode{}`一样
	root.Left.Right = createNode(2) //功能同`root.right.left = &treeNode{}`一样

	fmt.Println(root)             //{3 0xc00000c0a0 0xc00000c0c0}
	fmt.Println(*root.Right)      //{5 0xc00000c0e0 <nil>}
	fmt.Println(*root.Right.Left) //{0 <nil> <nil>} 无论地址还是结构本身，一律使用.来访问成员
	fmt.Println(*root.Left.Right) //{2 <nil> <nil>}

	/**
	 * 最后形成的树状结构如下：
			3
	      /   \
		0		5
	     \	   /
	      2   0
	*/

	//调用struce为treeNode定义的print()方法
	root.Print()     //输出3
	printValue(root) //输出3

	//修改struct为treeNode的value
	root.Right.Left.SetValue(99)     //输出3
	root.Right.Left.Print()          //输出0，无效修改
	root.Right.Left.SetTrueValue(99) //输出3
	root.Right.Left.Print()          //输出99，有效修改

	//输出中序遍历结果
	root.Traverse() //0 2 3 99 5

	//声明 struct型 数组，数组中每一个value表示一个struct
	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}

	fmt.Println(nodes) //[{3 <nil> <nil>} {0 <nil> <nil>} {6 <nil> 0xc00000c080}]
}

/**
 * 创建节点，返回一个 struct 结构为 treeNode 指针类型
 */
func createNode(num int) *treeNode {
	return &treeNode{value: num} //返回一个struct的引用
}
