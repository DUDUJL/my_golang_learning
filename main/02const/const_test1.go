package main

import "fmt"

// const用来定义枚举类型
// const (
// 	// 可以在const()添加一个关键字 iota， 每行的iota都会累加1，第一行的iota的默认值是0
// 	BEIJING = 0
// 	SHANGHAI = 1
// 	SHENZHEN = 2
// )

const (
	BEIJING = 10*iota //0
	SHANGHAI	   // 1
	SHENZHEN       // 2
)

const (
	a, b = iota+1, iota+2
	c, d
	e, f

	g, h = iota*2, iota*3
	i, j
)


func main()  {
	// 常量（只读属性）
	const length int = 10
	fmt.Println("length: ", length)

	// 
	fmt.Println("BEIJING = ", BEIJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)
	fmt.Println("SHENZHEN = ", SHENZHEN)
}