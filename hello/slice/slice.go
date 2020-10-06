package main

import (
	"fmt"
	"math"
	"sort"
)

func main()  {
	// a := [6]int {1,2,3,4,5,6}
	// b := a[1:4]
	// fmt.Printf("slice b = %v\n", b)
	// MakeSlice()
	// ModifySlice()
	// EmptySlice()
	// CopySlice()
	s := []float64{5.2, -1.3, 0.7, -3.8, 2.6} // unsorted
	sort.Float64s(s)
	fmt.Println(s)

	s = []float64{math.Inf(1), math.NaN(), math.Inf(-1), 0.0} // unsorted
	sort.Float64s(s)
	fmt.Println(s)

}

func MakeSlice()  {
	 var a []int
	 a = make([]int, 20, 40)
	 a[0] = 100
}

/*
切片是对数组的引用，切片指向的是数组的地址，切片做了修改相应的数组也会做出改变
 */
func ModifySlice()  {
	arr := [8]int {1,6, 3, 5, 6, 6}
	sli1 := arr[2:5]
	sli2 := arr[:]
	fmt.Printf("modify before arr = %v\n", arr)
	fmt.Printf("modify before sli1 = %v\n", sli1)
	fmt.Printf("modify before sli2 = %v\n", sli2)
	/*for i := range  sli1{
		sli1[i]++
	}*/
	sli1[0] = 100
	sli2[1] = 200
	fmt.Printf("modify after sli1 = %v\n", sli1)
	fmt.Printf("modify after sli2 = %v\n", sli2)
	fmt.Printf("modify after arr1 = %v\n", arr)
}

func EmptySlice()  {
	var slice []int
	fmt.Printf("slice addres = %p, len = %d, cap = %d", slice, len(slice), cap(slice))
}

func CopySlice()  {
	slice1 := []int {1,2,3,4}
	slice2 := []int {100,233,313,313,13,1}
	fmt.Printf("copy before slice1 = %v\n", slice1)
	copy(slice1, slice2)
	fmt.Printf("copy after slice1 = %v\n", slice1)
}