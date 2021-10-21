package main

import "fmt"

/**
	map的定义方式
*/

func main() {

	var myMap1 map[string]string // 括号内表示key类型 括号外表示value类型
	if myMap1 == nil {
		fmt.Println("myMap1 是一个空map")
	}

	// 在使用map前 先为map分配数据空间 10这个参数相当于前面slice的容量
	myMap1 = make(map[string]string, 10)
	myMap1["one"] = "java"
	myMap1["two"] = "python"
	myMap1["three"] = "go"
	fmt.Println(myMap1) // 打印出来是不讲顺序的

	// 第二种声明方式
	myMap2 := make(map[int]string)
	myMap2[1] = "java"
	myMap2[2] = "python"
	myMap2[3] = "go"
	fmt.Println(myMap2)

	// 第三种声明方式
	myMap3 := map[string]string{
		"one": "php",
		"two": "python",
	} // 注意这里最后一行要有逗号
	fmt.Println(myMap3)
}
