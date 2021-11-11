package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (this *User) Call() {
	fmt.Println("user is called ..")
	fmt.Printf("%v\n", this)
}

func main() {
	user := User{1, "zhang3", 18}
	user.Call()
	DoFiledAndMethod(user)
}

// 通过type获取里面的字段
// 1. 获取interface的reflect.Type，通过Type得到NumField，进行遍历
// 2. 得到每个field，数据类型
// 3. 通过field有一个interface()方法得到对应的value
func DoFiledAndMethod(input interface{}) {
	// 获取type和value
	inputType := reflect.TypeOf(input)
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputType is: ", inputType.Name())
	fmt.Println("inputValue is: ", inputValue)

	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Println(field)
		fmt.Println(value)
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	// 通过type 获取里面的方法
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("hahaha%s: %v\n", m.Name, m.Type)
	}
}
