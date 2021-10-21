package main

import "fmt"

func main() {
	var a string
	a = "aceld" // pair<static type:string, value: "aceld">

	var allType interface{} // pair<>
	allType = a             // par<type: string, value: "aceld">

	str, _ := allType.(string)
	fmt.Println(str)
}
