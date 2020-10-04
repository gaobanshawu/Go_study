package main

import (
	"fmt"
	"math"
)

var u uint8 = 255

func main() {
	fmt.Println(math.MaxInt8)
	fmt.Println(u, u+254, u*u)
	// x := -12
	// z := ^x
	// fmt.Println(x, z)
	// y := 1
	// yy := y << 5
	// fmt.Println(y, yy)
	// f := 1e100
	// o := int(f)
	i := 0777
	fmt.Printf("%d, %[1]o, %#[1]o\n", i)
	ii := 0xfff
	fmt.Printf("%d, %[1]o, %[1]x, %#[1]x\n", ii)
}
