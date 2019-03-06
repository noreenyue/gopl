package main

import "fmt"

func main() {
	// 赋值的左右必须具有相同的参数类型，==和!=也同样。
	v := 0
	s := "1"
	// v = s
	// if v != s {}
	fmt.Printf("v=%d, s=:%s\n", v, s)
}

// 最大公约数
func gcd(x, y int) int{
	for y!=0 {
		x, y = y, x%y
	}
	return x
}

// 斐波那契数列
// 0, 1, 1, 2, 3, 5, 8 ...
func fib(n int) int {
	x, y := 0, 1
	if 0 == n {
		return x
	}
	for i:=0; i<n; i++ {
		x, y = y, x+y
	}
	return x
}