package main

import "fmt"

// interface 本质是一个指针
type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

// 一个具体的
type Cat struct {
	color string
}

// 这里必须要重写接口的所有方法
func (this *Cat) Sleep() {
	fmt.Println("Cat is sleep")
}

func (this *Cat) GetColor() string {
	return this.color
}
func (this *Cat) GetType() string {
	return "Cat"
}

func showAnimal(animal AnimalIF) {
	animal.Sleep() // 多态
	fmt.Println("color =", animal.GetColor())
}

func main() {
	var animal AnimalIF // 接口的数据类型 父类指针
	animal = &Cat{"Green"}

	animal.Sleep() // 调用的就是Cat的Sleep方法 多态现象

	showAnimal(animal)

}
