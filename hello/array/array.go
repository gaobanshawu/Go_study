package main

import (
	"fmt"
)

func main()  {
	// fmt.Println("jiushiwo")
	// a := [...]int {0:100, 3:300}
	// for index, value := range a  {
	// 	fmt.Printf("index = %d, value = %v\n", index, value)
	// }

	// a := [3]int {10, 20, 30}
	// b := a
	// b[0] = 80
	// fmt.Printf("a = %v, b = %v", a, b)

	a := [6]int {1, 3, 5, 7, 8, 7}
	fmt.Printf("a的和为： %d\n", AddArray(a))
	Sum8(a, 8)
}

func AddArray(a [6]int) int {
	sum := 0
	for i := 0; i < len(a); i++{
		sum += a[i]
	}
	return sum
}

func Sum8(a [6]int, sum int) {
	for i := 0; i < len(a) ; i++{
		for j := i+1; j < len(a); j++{
			if a[i] + a[j] == sum {
				fmt.Printf("(%d, %d)\n", i, j)
			}
		}
	}
}
