package main

import (
	"fmt"
	"hello/package_study"
	"strings"
)

func TestStringStudy() {
	// 练习二
	var zh_str string = "上海自来水来自海上"
	var rune_zh_str []rune = []rune(zh_str)

	for i := 0; i <= len(rune_zh_str)/2; i++ {
		if rune_zh_str[i] == rune_zh_str[len(rune_zh_str)-i-1] {
			continue
		} else {
			fmt.Printf("字符串不是回文")
		}
	}
	fmt.Printf("字符串是回文")

	// 练习一
	var en_str string = "jiushiwo"
	var rune_en_str []rune = []rune(en_str)

	for i := 0; i <= len(rune_en_str)/2; i++ {
		empty := rune_en_str[len(rune_en_str)-i-1]
		rune_en_str[len(rune_en_str)-i-1] = rune_en_str[i]
		rune_en_str[i] = empty
	}
	versa_en_str := string(rune_en_str)
	fmt.Printf("%s\n", versa_en_str)
}

func TestString() {
	var str1 string = "hello world"
	fmt.Printf("%s\n", str1)

	// 去取字符串的字符长度
	var str1_lens int = len(str1)
	fmt.Printf("str1_lens = %d\n", str1_lens)

	// 判断是否存在子字符串
	var isin bool = strings.Contains(str1, "hell")
	fmt.Printf("%v\n", isin)

	// 判断字符串以什么开头
	var isj bool = strings.HasPrefix(str1, "j")

	// 判断字符串以什么结尾
	var isd bool = strings.HasSuffix(str1, "d")
	fmt.Printf("%v\n", isj)
	fmt.Printf("%v\n", isd)

	// `   ` 打印原始字符串
	var shiju string = `
				空山新雨后，天气晚来秋。
				明月松间照，清泉石上流。
				竹喧归浣女，莲动下渔舟。
				随意春芳歇，王孙自可留。
`
	fmt.Printf("%s\n", shiju)

	// 字符串拼接
	var str2 string = "jiu shi wo zen mo yang"
	var str string = fmt.Sprintf("%s%s", str1, str2)
	fmt.Printf("%s\n", str)

	// 分割字符串 用什么样的格式区分割字符串spilt,返回一个切片即列表
	yuju := strings.Split(str2, " ")
	fmt.Printf("%s\n", yuju[0])

	// 判断子字符串第一次出现的位置
	first_index := strings.Index(str2, "z")
	fmt.Printf("%d\n", first_index)

	// 判断子字符串最后出现的位置
	last_index := strings.LastIndex(str2, "n")
	fmt.Printf("%d\n", last_index)

	// 将数组元素以某种符号连接  返回字符串
	strarr := []string{"heitiejingjie", "jiushiwo", "yinhao"}
	str3 := strings.Join(strarr, "-")
	fmt.Printf("%s\n", str3)

	// 修改字符串的值要把字符串先转成byte切片类型， 在修改切片的值，修改完成转回字符串
	fmt.Printf("before modify str2 = %s\n", str2)
	var sectionbyte []byte = []byte(str2)
	sectionbyte[1] = 'h'
	str2 = string(sectionbyte)
	fmt.Printf("after modify str2 = %s\n", str2)

	// rune类型用来表示utf-8的类型，一个rune字符由一个或多个字节组成
	var runes rune = '就'
	fmt.Printf("%c\n", runes)
}

func PackageVisit() {
	fmt.Printf("%s\n", package_study.A)
}
func main() {
	TestString()
	PackageVisit()
	TestStringStudy()
}
