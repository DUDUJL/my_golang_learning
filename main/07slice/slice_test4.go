package main

import "fmt"

/**
slice切片的追加与截取
*/

func main() {

	var numbers = make([]int, 3, 5) // 长度为3 容量为5 （开辟了5块空间，但合法元素只有3个） 尾部指针在第三个元素的位置
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	// 向numbers追加一个元素1
	numbers = append(numbers, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	// 向numbers追加一个元素2
	numbers = append(numbers, 2)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	// 容量满了 再去append
	numbers = append(numbers, 3) // 这时会再一次性多开辟5块存储空间
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	s := []int{1, 2, 3} // len = 3, cap = 3
	s1 := s[0:2]
	fmt.Println(s1)

	s2 := s[:2]
	fmt.Println(s2)

	s3 := s[2:]
	fmt.Println(s3)

	// 修改s1的第0个元素
	s1[0] = 111
	fmt.Println(s[0]) // 可知 s s1 s2 s3 指向的都是同一个地址

	// copy 可以将底层数组的slice一起进行拷贝
	s4 := make([]int, 3) // [0, 0, 0]

	copy(s4, s) // 将s中的值拷贝到s4中

}
