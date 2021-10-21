package main

import "fmt"

func main() {
	// defer关键字相当于Java中的final 是压栈的形式 先进栈的后出栈
	// defer的调用还要在return关键字之后
	defer fmt.Println("main end1")
	defer fmt.Println("main end2")

	fmt.Println("hello go 1")
	fmt.Println("hello go 2")

}
