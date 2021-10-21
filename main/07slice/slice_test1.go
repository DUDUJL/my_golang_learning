package main

import "fmt"

// [4]int 与 [10]int 是两种数据类型
func printArray(myArray [4]int) {
	// 这是一个值拷贝
	for index, value := range myArray {
		fmt.Println(index, value)
	}
}

func main() {
	// 一个固定长度的数组
	var myArray1 [10]int
	// 数组的遍历
	for i := 0; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])
	}

	myArray2 := [10]int{1, 2, 3, 4} // 这样赋值 前4个参数有值 其他都是默认值0
	// range的遍历写法
	for index, value := range myArray2 {
		fmt.Println(index, value)
	}

	myArray3 := [4]int{1, 2, 3, 4}

	printArray(myArray3)
}
