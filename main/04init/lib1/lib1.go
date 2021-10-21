package lib1

import "fmt"

// 函数名称开头大写 表示对外开放； 小写则只能内部使用
func Lib1Test() {
	fmt.Println("lib1Test...")
}

func init() {
	fmt.Println("lib1 init ...")
}
