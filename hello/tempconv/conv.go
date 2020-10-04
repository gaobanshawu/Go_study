package tempconv

// CtoF ... 将摄氏度转换为华氏度
func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FtoC ... 将华氏度转换为摄氏度
func FtoC(f Fahrenheit) Celsius {
	return Celsius(f-32) * 5 / 9
}
