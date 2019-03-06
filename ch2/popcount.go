package main

import (
	"fmt"
)


var pc[256]byte
// 用于初始化0-255的bit数，key为数值，val为bit数
func init() {
	// i 是pc的索引
	// 由于是递增，没增加一位只需要用上一位计算结果，加上最低位的奇偶性
	for i := range pc {
		fmt.Printf(" i=%v ---> byte(%d&1)=%v ", i, i, byte(i&1))
		fmt.Printf("%d/2=%d, pc[%d]=%v ", i, i/2, i/2, pc[i/2])
		fmt.Printf("pc[%d]=%v\n", i, pc[i/2]+byte(i&1))
		pc[i] = pc[i/2] + byte(i&1)
		fmt.Printf("pc=%v\n", pc)
	}
}

func main() {
	x := 8 + 4 << 8
	fmt.Printf("x=%d\n", x)

	/*
		byte类型是一个uint8，高位抹零
		byte(x>>(0*8)) 取低8位
		byte(x>>(1*8)) 右移8位后，取低8位
		pc[byte(x>>(0*8))] 对每一节8位，做count 1bit计算
	  */
	fmt.Printf("x>>(0*8) %v, %v\n", x>>(0*8), byte(x>>(0*8)))
	fmt.Printf("x>>(1*8) %v, %v\n", x>>(1*8), byte(x>>(1*8)))

	fmt.Println(PopCount(uint64(x)))
}

func PopCount(x uint64) int {

	ret := 0
	for i:=uint(0); i<8; i++ {
		ret += int(pc[byte(x>>(i*8))])
	}
	return ret
	//return int(pc[byte(x>>(0*8))] +
	//	pc[byte(x>>(1*8))] +
	//	pc[byte(x>>(2*8))] +
	//	pc[byte(x>>(3*8))] +
	//	pc[byte(x>>(4*8))] +
	//	pc[byte(x>>(5*8))] +
	//	pc[byte(x>>(6*8))] +
	//	pc[byte(x>>(7*8))])
}