package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat() ...")
}

func (this *Human) Walk() {
	fmt.Println("Human.Walk() ...")
}

type Superman struct {
	Human // 表示Superman类继承了Human类的方法
	level int
}

//重定义父类方法Eat()
func (this *Superman) Eat() {
	fmt.Println("Superman.Eat() ...")
}

// 子类的新方法
func (this *Superman) Fly() {
	fmt.Println("Superman.Fly() ...")
}

func main() {
	h := Human{"zhang3", "male"}

	h.Eat()
	h.Walk()

	// 定义一个子类对象
	s := Superman{Human{"li4", "female"}, 1}

	s.Walk()
	s.Eat()
	s.Fly()

	// 另外一种写法
	var s_ Superman
	s_.name = "wang5"
	s_.sex = "male"
	s_.level = 2

	
}
