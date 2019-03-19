package main

import (
	"fmt"
)

/**
 反转Slice
 */
func main() {
	a := [...]int{ 1, 2, 3, 4, 5}
	s := a[:]
	fmt.Println(s)
	reserve(s)
	fmt.Println(s)

	// slice的初始化方法与数组略有不同
	s2 := []int{1,2,3,4,5}
	// 向左推n元素
	leftN(2, s2)
	fmt.Println(s2)
}

func reserve(s []int){
	for i, j:=0, len(s)-1; i<j; i,j = i+1,j-1  {
		s[i], s[j] = s[j], s[i]
	}
}


/**
 向左反转n个元素
 三次反转
  */
func leftN(pos int, s []int) {
	reserve(s[:pos])
	reserve(s[pos:])
	reserve(s)
}