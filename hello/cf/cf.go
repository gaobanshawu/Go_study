package main

import (
	"fmt"
	"hello/tempconv"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[0:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprint(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s == %s, %s == %s", f, tempconv.FtoC(f), c, tempconv.CtoF(c))
	}
}
