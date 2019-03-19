package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	/**
	%x 以十六进制格式打印， %t布尔类型，%T变量类型
	 */
	fmt.Printf("`c1` is %x\n`c2` is %x\n`c1==c2` is %t\n`type of `c` is %T\n", c1, c2, c1==c2, c1)

	count := make(map[byte]int)
	countBit(c1, count)
	countBit(c2, count)
	fmt.Println(count)

	input := bufio.NewScanner(os.Stdout)
	for input.Scan(){
		shaCode("x",input.Text());
	}
}

// 给byte数组清零
func zero(ptr * [32]byte) {
	// 逐个赋值
	for i := range ptr{
		ptr[i] = 0
	}
	// 整体赋值
	*ptr = [32]byte{}
}

func countBit(ptr [32]byte, count map[byte] int){
	for i := range ptr {
		count[ptr[i]] ++
	}
}

func shaCode(s string, shaType string)  {
	switch shaType {
	case "Sum384":
		fmt.Println(sha512.Sum384([]byte(s)))
	case "Sum512":
		fmt.Println(sha512.Sum512([]byte(s)))
	default:
		fmt.Println(sha256.Sum256([]byte(s)))
	}
}