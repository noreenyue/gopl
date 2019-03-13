package main

import (
	"fmt"
	"math/cmplx"
)

/**
 复数
*/
func main() {
	var x complex128 = complex(1, 1)	// 1+1i
	var y complex128 = complex(2, 1)	// 2+1i

	fmt.Println(x*y)	// (1+3i)
	fmt.Println(real(x*y))	// real 实数部分 1
	fmt.Println(imag(x*y))	// imag 虚数部分 3

	fmt.Println(1i*2i)	// (-2+0i)	1^2=-2

	m := 1 + 2i
	n := 3 + 4i
	fmt.Println(m, n)
	fmt.Println(cmplx.Sqrt(-1))
}
