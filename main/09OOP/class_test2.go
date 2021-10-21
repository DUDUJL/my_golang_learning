package main

import "fmt"

// 一个普通的方法 如果第一个字母大写 表示可以被外部访问
// 一个类名首字母大写 表示该属性是可以被外部访问的 否则的话只能类的内部访问
type Hero struct {
	Name  string // 公有属性
	Age   int
	Level int
	des   string // 私有属性
}

// 注意：对象前面加 * 号
func (this *Hero) GetName() string {
	return this.Name
}

func (this *Hero) SetName(newName string) {
	this.Name = newName
}

func main() {
	hero := Hero{Name: "zhang3", Age: 100, Level: 1}

	hero.GetName()
	hero.SetName("li4")
	fmt.Printf("%v", hero)
}
