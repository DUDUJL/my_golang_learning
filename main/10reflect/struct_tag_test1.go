package main

import (
	"fmt"
	"reflect"
)

// 带tag的结构体
type resume struct {
	Name string `info:"name" doc:"我的名字"`
	Sex  string `info:"sex"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()
	for i := 0; i < t.NumField(); i++ {
		tagstring := t.Field(i).Tag.Get("info")
		fmt.Println(tagstring)
	}
}

func main() {
	var re resume
	findTag(&re)
}
