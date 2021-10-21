package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// tty: pair<type:*os.File, value: "/dev/tty"文件描述符>
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)

	if err != nil {
		fmt.Println("open file error", err)
		return
	}

	// r pair<type: , value: >
	var r io.Reader // 这Reader就是一个interface
	r = tty         // r pair<type:*os.File, value:"/dev/tty"文件描述符>

	// w pair<type: , value: >
	var w io.Writer
	w = r.(io.Writer) // w pair<type:*os.File, value:"/dev/tty"文件描述符>

	w.Write([]byte("HELLO THIS is A TEST!!!\n"))
}
