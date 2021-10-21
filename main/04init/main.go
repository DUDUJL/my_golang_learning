package main

// 导入 但不使用 编译器会报错，加'_'表示匿名导入 这样编译器不报错 又执行了init方法
import (
	"awesomeProject/main/04init/lib1"
	_ "awesomeProject/main/04init/lib2"      // 匿名导入 无法使用当前包的方法，但是会执行init函数
	mylib3 "awesomeProject/main/04init/lib3" // 取一别名
	. "fmt"                                  // 将fmt中的所有方法 导入到当前的作用域中，不需要 fmt.xxx 来调用
)

func main() {
	lib1.Lib1Test()
	mylib3.Lib3Test()
	Println("hahaha")
}
