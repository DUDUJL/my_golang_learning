package main

import (
	"fmt"
	"reflect"
)

/**
ValueOf()和TypeOf()
ValueOf()用来获取输入参数接口中的数据的值，接口为空则返回0
TypeOf()用来动态获取输入参数接口中值的类型，如果接口为空则返回nil
*/

func reflectNum(arg interface{}) {
	fmt.Println("type: ", reflect.TypeOf(arg))
	fmt.Println("value: ", reflect.ValueOf(arg))
}

func main() {
	var num float64 = 1.2345
	reflectNum(num)
}
