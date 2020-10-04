package main

func main() {
	// 	const freezing, boilingF = 32.0, 212.0
	// 	fmt.Printf("%g℉ = %g℃\n", freezing, fToc(boilingF))
	// 	fmt.Printf("%g℉ = %g℃\n",  boilingF, fToc(freezing))
	// var j, k, l  int
	// fmt.Printf("%d, %d, %d", j, k, l)
	/*
		f, err := os.Open(name)
		if err != nil{
			return err
		}

		f.close()
	*/
	// i, n := 8, 9
	//
	// i, n := 8, 9
	/*
		i := 1
		p := &i
		println(p)
		println(*p)
		println("------------------------------------")
		*p = 3
		println(p)
		println(*p)
		var x, y, z int
		println(&x == &x, &x == &y, &z == nil)
		println(&x)
		println(&y)
		println(&z)

		v := 1
		incr(&v)
		println(incr(&v))
	*/
	// 使用new函数，将创建一个基本类型的匿名变量，返回该变量的内存地址,初始值为基本类型的零值
	/*
		p := new(int)
		p1 := new(string)
		p2 := new(float64)
		println(p)
		println(*p)
		println("---------------------")
		println(p1)
		println(*p1)
		println("---------------------")
		println(p2)
		println(*p2)
	*/

	println(fib(12))

	// f, err := os.Open("../time/time_study.go")
	// if f != nil{
	// 	return err
	// }

}

// func fToc(f float64) float64 {
// 	return (f - 32) * 5 / 9

// func incr(p *int) int {
// 	*p++
// 	return *p
// }

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
