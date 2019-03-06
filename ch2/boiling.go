package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	// 函数内的局部变量，必须在被调用之前声明
	var c = ( f - 32 ) * 5 / 9

	// %f, %g, %e float类型
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
}

// 包一级变量，可以再任意位置声明
//const boilingF = 212.0