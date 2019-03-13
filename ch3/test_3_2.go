package main
/**
 浮点型
*/
import (
	"fmt"
	"math"
)

func main() {
	// float32最大值
	fmt.Println(math.MaxFloat32, math.MinInt32)

	/*
		float32的有效bit位为23个，其他bit用于表示指数和符号
		超过23位出现误差
	*/
	var f float32 = 1 << 24
	fmt.Println(f, f+1, f == f+1)

	const e = 2.71828


	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d, e^%[1]d = %8.3f\n", x, math.Exp(float64(x)))
	}

	/*
	 * 正无穷+Inf 负无穷-Inf
	 * 无效操作NaN, 比如0/0、Sqrt(-1)等
	 * math.IsNaN用于判断是否为NaN
	 */
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z, math.IsNaN(z/z))

	// NaN和任何值都不相等
	man := math.NaN()
	fmt.Println(man==man, man!=man, man>man, man<man)	// false, true, false, false


}
