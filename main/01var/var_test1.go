package main

import "fmt"

// 声明全局变量 方法123是可以的
var gA int = 100
var gB = 200

// 而:=只能用在函数体内来声明
// gC := 300

func main()  {
	// 声明变量 默认值是0
	var a int
	fmt.Println("a = ", a)

	//声明变量 初始化一个值
	var b int = 100
	fmt.Println("b = ", b)

	// 初始化的时候省去数据类型 通过值自动匹配当前变量的数据类型
	var c = 100
	fmt.Println("c = ", c)
	fmt.Printf("type of c is %T\n", c)

	var s string = "abcd"
	fmt.Printf("s = %s, type of s is %T\n", s, s)

	// 最常用的方法 省去var关键字 直接自动匹配
	d := 100 // mod表示即初始化 又赋值
	fmt.Printf("type of d is %T\n", d)
	e := 3.14
	fmt.Printf("type of e is %T\n", e)

	// 同时声明多个变量
	var xx, yy int = 100, 200 // 同种数据类型
	var kk, ll = 100, "ll" // 不同数据类型
	fmt.Println(xx, yy, kk, ll)

	// 多行多变量声明
	var (
		vv int = 100
		jj bool = true
	)
	fmt.Println(vv, jj)
}