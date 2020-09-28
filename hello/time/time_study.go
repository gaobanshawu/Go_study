package main

import (
	"fmt"
	"time"
)

// 时间的一些练习
func time_study() {
	// 获取当前时间
	now_time := time.Now()
	fmt.Printf("%v\n", now_time)
	// 获取当前的时间详情
	year := now_time.Year()
	month := now_time.Month()
	day := now_time.Day()
	hour := now_time.Hour()
	minute := now_time.Minute()
	second := now_time.Second()
	fmt.Printf("year = %v, month = %v, day = %v, hours = %v, minute = %v, second = %v\n", year, month, day, hour,
		minute, second)
	// 获取当前时间戳
	fmt.Printf("%v\n", now_time.Unix())

}
func main() {
	time_study()
}
