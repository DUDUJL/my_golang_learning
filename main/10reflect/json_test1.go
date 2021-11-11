package main

import (
	"encoding/json"
	"fmt"
)

/**
结构体标签在json中的应用
*/

// 要转成json必须要给他定标签
type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"price"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"hahaha", 2000, 10, []string{"zhang3", "li4"}}

	// 编码的过程 结构体--->json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error:", err)
		return
	}
	fmt.Printf("jsonStr = %s\n", jsonStr)
}
