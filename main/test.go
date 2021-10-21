package main

import "fmt"
import "time"

func main() { // 函数的“{”一定和函数名在同一行，否则要报错
	// golang中“;”可加可不加
	const name string = "dujianlin"
	fmt.Println(name)
	time.Sleep(1 * time.second)
}