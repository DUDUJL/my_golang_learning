package main

import "fmt"

/**
一个关于指针的经典例子
*/
func swap(a *int, b *int) {
	c := *a
	*a = *b
	*b = c
}

func main() {
	var a int = 10
	var b int = 20

	swap(&a, &b)

	fmt.Println("a =", a, " b =", b)
	fmt.Println(&a)
}
