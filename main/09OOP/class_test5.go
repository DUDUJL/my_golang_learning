package main

import "fmt"

/**
通用万能类型
interface{} 空接口
int、string、float32、float64、struct 都实现了interface{}

可以用interface{}类型 引用 任意的数据类型
*/

func myFunc(arg interface{}) {
	fmt.Println("myFunc is called ...")
	fmt.Println(arg)

	// interface 该如何区分 此时引用的数据类型到底是什么
	// 给interface提供断言
	value, ok := arg.(string) // 判断arg是否为string

	if !ok {
		fmt.Println("arg is not string")
	} else {
		fmt.Println("arg is string type, value =", value)
		fmt.Printf("value type is %T\n", value)
	}

}

type Book_ struct {
	auth string
}

func main() {
	book := Book_{"Golang"}

	myFunc(book)
	myFunc("abc")
	myFunc(3.14)
}
