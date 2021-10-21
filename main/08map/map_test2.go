package main

import "fmt"

/**
map的使用方式
*/

func printMap(cityMap map[string]string) {
	// 这是一个引用传递 这里的cityMap和外面的指向的是同一块内存区域
	// 遍历
	for key, value := range cityMap {
		fmt.Println("key = ", key)
		fmt.Println("value = ", value)
	}
}

func main() {
	cityMap := make(map[string]string)

	// 添加
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"

	printMap(cityMap)

	// 删除
	delete(cityMap, "China")

	// 修改
	cityMap["USA"] = "DC"

}
