package main

import "fmt"

// 声明一种新的数据类型 myint，实际上就是int的一个别名
type myint int

// 定义一个结构体
type Book struct {
	title string
	auth  string
}

func changeBook1(book Book) {
	// 传递一个book的副本
	book.auth = "666"
}

func changeBook2(book *Book) {
	book.auth = "777"
}

func main() {
	var a myint = 10
	fmt.Println("a = ", a)
	fmt.Printf("type for a is %T\n", a)
	var book1 Book
	book1.title = "Golang"
	book1.auth = "zhang3"
	fmt.Printf("%v\n", book1)
	changeBook1(book1)
	fmt.Printf("%v\n", book1) // 就发现改不了值

	changeBook2(&book1)
	fmt.Printf("%v\n", book1)
}
