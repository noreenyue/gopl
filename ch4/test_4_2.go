package main

import (
	"bytes"
	"fmt"
)

/**
 Slice 切片
 三部分构成：指针、长度、容量
 */
func main() {
	months := [...]string{
		1 : "1月", 2 : "2月", 3 : "3月", 4 : "4月",
		5 : "5月", 6 : "6月", 7 : "7月", 8 : "8月",
		9 : "9月", 10 : "10月", 11 : "11月", 12 : "12月",
	}

	Q2 := months[4:7]
	summer := months[6:9]
	// slice之间可以共享底层数据，并且引用的数据区间可以重叠
	for _, monthOfQ := range Q2 {
		for _, monthOfSummer := range summer {
			if monthOfQ == monthOfSummer {
				fmt.Printf("%s appears in both\n", monthOfQ)
			}
		}
	}

	fmt.Println(Q2[:5])		// 超出长度(3)，自动扩展
	//fmt.Println(Q2[:14])	// 超出容量(13)，导致panic


	// slice之间不能相互比较
	//fmt.Println(s == s2)	// compile error
	aByte := []byte("string")
	bos1 := aByte[:2]
	bos2 := aByte[2:5]
	fmt.Println(bos1, bos2, bytes.Equal(bos1, bos2))

	aIntegers := [...]int{6, 7, 8, 6, 7}
	x := aIntegers[:2]
	y := aIntegers[3:5]
	fmt.Println(x, y, equal(x, y))
}

func equal(x, y[] int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
