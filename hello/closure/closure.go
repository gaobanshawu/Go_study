package main

import (
	"fmt"
	"strings"
)

func main(){
	var f = Adder()
	fmt.Printf("-----%v\n", f(1))
	fmt.Printf("-----%v\n", f(200))
	fmt.Printf("-----%v\n", f(30))

	func1 := MakeSuffix(".bmp")
	func2 := MakeSuffix(".jpg")
	func1("jiushiwo")
	func2("heitiejingjie")

	calc1, calc2 := calc(10)
	fmt.Printf("calc1 = %d, calc2 = %d\n", calc1(1), calc2(2))
	fmt.Printf("calc1 = %d, calc2 = %d\n", calc1(1), calc2(2))
	fmt.Printf("calc1 = %d, calc2 = %d\n", calc1(1), calc2(2))
	fmt.Printf("calc1 = %d, calc2 = %d\n", calc1(1), calc2(2))

	var isort[8]int = [8]int{1, 7, 2, 4, 0, 4, 8, 3}
	// fmt.Println(insert_sort(isort))
	// fmt.Println(select_sort(isort))
	fmt.Println(bullte_sort(isort))
}

func Adder() func(int) int {
	var x int
	return func(i int) int {
		x += i
		return x
	}
}

func MakeSuffix(suffix string) func(string) string  {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix){
			return name + suffix
		}
		return name
	}
}

func calc(base int) (func (int) int, func(int) int ) {
	add := func(x int) int {
		base += x
		return base
	}
	sub := func(y int) int {
		base -= y
		return base
	}
	return add, sub
}

func insert_sort(ii [8]int)[8]int {
	for i := 1; i < len(ii); i++{
		for j := i; j > 0; j--{
			if ii[j] > ii[j-1] {
				ii[j], ii[j-1] = ii[j-1], ii[j]
			}else{
				break
			}
		}
	}
	return ii
}

func select_sort(ii [8]int) [8]int {
	for i := 0; i < len(ii); i++ {
		for j := i + 1; j < len(ii); j++ {
			if ii[j] > ii[i] {
				ii[i], ii[j] = ii[j], ii[i]
			}
		}
	}
	return ii
}

func bullte_sort(ii [8]int) [8]int {
	for i := len(ii)-1; i >= 0; i-- {
		for j := i-1; j >= 0; j-- {
			if ii[i] < ii[j] {
				ii[i], ii[j] = ii[j], ii[i]
			}
		}
	}
	return ii
}