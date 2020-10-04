package main

import (
	"fmt"
)

func PrintHello() {
	fmt.Printf("Hello world")
	// _下划线可忽略参数
}

// Add ...参数类型  表示可变参数
func Add(a ...int) int {

	return 0
}

// Sum defer语句在函数返回之前执行, 适合用来做资源的释放
func Sum() {
	defer fmt.Printf("函数将要返回了哟！\n")
	fmt.Printf("zheshiwo\n")
	fmt.Printf("heitiejingjie\n")
}

// 匿名函数  没有函数名字

func main() {
	PrintHello()
	Sum()
}
