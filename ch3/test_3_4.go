package main

import "fmt"

func main() {
	s := ""
	// 左边判断失败，右边不再执行。安全
	if s != "" && s[0] != 'x' {
		fmt.Println("s[0]==x")
	}
	// fmt.Println(s[0])	// 越界导致panic

	s = "1"
	c := s[0]
	// && 优先级高于 ||
	if 'a' <= c && c <= 'z' ||
		'A' <= c && c <= 'Z' ||
		'0' <= c && c <= '9' {
		fmt.Println(c)
	}

	// int型 和 bool型 不能隐式转换
	//fmt.Println( 1 == true) 	// compile error
}
