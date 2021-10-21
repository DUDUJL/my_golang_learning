package main

import "fmt"

/**
slice切片的四种声明定义方式
*/

func main() {
	// 声明一个slice切片，并且初始化，默认值是1 2 3，长度为3
	slice1 := []int{1, 2, 3}
	fmt.Printf("len = %d, slice1 = %v\n", len(slice1), slice1) // %v 打印详细信息 任何类型

	// 声明slice2是一个切片，但是没有给它分配空间
	var slice2 []int
	fmt.Printf("len = %d, slice2 = %v\n", len(slice2), slice2)

	// 分配3个空间 初始值0
	var slice3 []int = make([]int, 3) // slice3 := make([]int, 3) 通过:=推导出slice是一个切片
	fmt.Printf("len = %d, slice3 = %v\n", len(slice3), slice3)

	// 判断一个slice是否为空
	if slice2 == nil {
		fmt.Println("slice2 是一个空切片")
	} else {
		fmt.Println("slice2 是有值的")
	}
}
