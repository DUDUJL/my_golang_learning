package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

// 具体类型
type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("Read a book")
}

func (this *Book) WriteBook() {
	fmt.Println("Write a book")
}

func main() {
	// b: pair<type:Book, value:book{}地址>
	b := &Book{}

	// r: pair<type: , value: >
	var r Reader
	r = b // r: pair<type: Book, value: book{}地址>

	r.ReadBook()

	var w Writer
	// r: pair<type: Book, value: book{}地址>
	w = r.(Writer) // 此处为什么会断言成功？ 因为 w r 的具体type是一样的

}
