package main

import "fmt"

func main() {
	/*
	 *	 1    1    1    1   1   1   1   1
	 * 128 + 64 + 32 + 16 + 8 + 4 + 2 + 1	 = 256
	 */
	var u uint8 = 255
	// 无符号类型u+1溢出(1|0000|0000)，高位截断后值为0
	fmt.Println(u, u+1, u*u)

	var i int8 = 127
	// 有符号类型 0111|1111，i+1为(1000|0000)，有符号最高位为负
	fmt.Println(i, i+1, i*i)


	var x uint8 = 1<<5 | 1<<1
	var y uint8 = 1<<2 | 1<<1
	fmt.Printf("%08b\n", x)	// 00100010
	fmt.Printf("%08b\n", y)	// 00000110
	fmt.Printf("%08b\n", x&y)		// 按位与	   遇0则0
	fmt.Printf("%08b\n", x|y)		// 按位或   遇1则1
	fmt.Printf("%08b\n", x^y)		// 按位异或 不同为1，相同为0
	fmt.Printf("%08b\n", x&^y)	// 按位置零 y中值为1的bit，x对应位置零

	// x : 00100010
	for i := uint8(0); i < 8; i++ {
		// 1左移i位使bit位与x重合: 1,5
		if x&(1<<i) != 0{
			fmt.Println(i)
		}
	}

	fmt.Printf("%08b\n", x<<1)	// 01000100
	fmt.Printf("%08b\n", x>>1)	// 00010001


	f := 3.141
	i2 := int(f)
	fmt.Println(f, i2)

	f = 1.99
	fmt.Println(int(f))

	f2 := 1e100
	fmt.Println(int(f2))

	o := 0666	// 6*64 + 6*8 + 1*6 = 438
	o2 := 0766
	/*
	    副词[1]: 表示再次使用第一个操作数
		副词#: 表示在用%o、%x、%X格式输出时，带上对应的0,0x,0X前缀
	 */
	fmt.Printf("%d %[1]o %d %#[2]o\n", o, o2)


	ascii := 'a'
	unicode := '国'
	newline := '\n'
	// %c打印字符，或者%q打印带单引号的字符
	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	fmt.Printf("%d %[1]c %[1]q\n", unicode)
	fmt.Printf("%d %[1]c %[1]q\n", newline)	// %c回车, %q打印字符
}
