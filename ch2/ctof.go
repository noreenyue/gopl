package main

import "fmt"

func main() {
	const freezingC, boilingC = 0, 100
	fmt.Printf("freezing point = %g°F or %d°C\n", ctof(freezingC), freezingC )
	fmt.Printf("boiling point = %g°F or %d°C\n", ctof(boilingC), boilingC )
}

/*
	函数定义格式：
	func 函数名 参数列表 返回值列表(optional) {
		// 函数体
	}
*/
func ctof(c int) float64{
	return float64(c) * 9 / 5 + 32
}