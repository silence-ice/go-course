package main

import (
	"fmt"
)

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {

	//最简单访问方式
	fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})                                  // 创建一个新的结构体
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407}) // 也可以使用 key => value 格式
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})                                       // 忽略的字段为 0 或 空

	var Book1 Books

	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	/* 打印 Book1 信息 */
	fmt.Printf("Book 1 title : %s\n", Book1.title)
	fmt.Printf("Book 1 author : %s\n", Book1.author)
	fmt.Printf("Book 1 subject : %s\n", Book1.subject)
	fmt.Printf("Book 1 book_id : %d\n", Book1.book_id)

	//使用函数 打印 Book1 信息
	printBook(Book1)

	/*
	 1. 结构体的指针类似于其他指针变量，格式如下：var struct_pointer *Books
	 2. 使用 struct_pointer = &Book1 传递变量地址
	*/
	pointPrintBook(&Book1)
}

func printBook(book Books) {
	book.title = "123" //修改title
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}

func pointPrintBook(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}
