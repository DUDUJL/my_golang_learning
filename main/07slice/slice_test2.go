package main

import "fmt"

// 形参数组长度不受限制
func printArray(myArray []int) {
	// 这个是引用传递
	// _ 表示匿名的变量
	for _, value := range myArray {
		fmt.Println(value)
	}
}

func main() {
	// 动态数组 slice
	myArray := []int{1, 2, 3, 4}
	fmt.Printf("the type of myArray is %T\n", myArray)
	printArray(myArray)
}
