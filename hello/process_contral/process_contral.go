package main

import (
	"fmt"
)

func if_process() {
	num := 13
	if num%2 == 0 {
		fmt.Printf("num = %d is even\n", num)
	} else {
		fmt.Printf("num = %d is odd\n", num)
	}
}
func swith_peocess() {
	// fallthrough 语句会穿透当前语句
	num := 10
	switch {
	case num >= 1 && num < 10:
		fmt.Printf("num >= 1 && num < 10\n")
	case num >= 10 && num < 20:
		fmt.Printf("num >= 10 && num < 20\n")
		fallthrough
	case num >= 20 && num < 30:
		fmt.Printf("num >= 20 && num < 30\n")
	}

}

func PrintMul() {
	for i := 1; i < 10; i++ {
		for ii := 1; ii <= i; ii++ {
			fmt.Printf("%d * %d = %d\t", ii, i, ii*i)
		}
		print("\n")
	}
}

func main() {
	if_process()
	swith_peocess()
	PrintMul()
}
