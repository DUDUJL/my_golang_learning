package main
import "fmt"

// 一个返回值的
func foo1(a string, b int) int {
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	c := 100
	return c
}

// 两个返回值的，匿名的
func foo2(a string, b int) (int, int) {
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	
	return 666, 777
}

// 两个返回值，有形参名的
// 这种情况下 r1 r2 是有初始化的 初始值是0 作用域空间是foo3内部
func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("----foo3----")
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	
	// 给有名称的返回值变量赋值
	r1 = 1000
	r2 = 2000

	return
}

// 两个返回值都是int
func foo4(a string, b int) (r1, r2 int) {
	fmt.Println("----foo4----")
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	
	// 给有名称的返回值变量赋值
	r1 = 1000
	r2 = 2000

	return
}

func main() {
	c := foo1("abc", 555)
	fmt.Println(c)

	ret1, ret2 := foo2("hahaha", 999)
	fmt.Println(ret1, ret2)

	ret1, ret2 = foo3("foo3", 333)
	fmt.Println(ret1, ret2)

	ret1, ret2 = foo3("foo4", 444)
	fmt.Println(ret1, ret2)
}