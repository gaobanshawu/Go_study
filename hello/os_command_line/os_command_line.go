package main

import (
	"fmt"
	"os"
	"strings"
)

// 获取外部输入的命令
func PrintOs() {
	// 第一种形式
	/*
		var s, sep string
		for i := 0; i < len(os.Args); i++ {
			s += sep + os.Args[i]
			sep = " "
		}
		fmt.Println(s)
	*/
	// 第二种形式
	s, sep := "", ""
	for index, arg := range os.Args[0:] {
		s += string(index)
		s += arg + sep
		sep = " "
	}
	fmt.Println(s)
	// 第三种形式
	fmt.Println(strings.Join(os.Args[0:], " "))
}

func main() {
	PrintOs()
}
