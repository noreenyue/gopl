package main

import "fmt"

/*
 数组:
  数字是由一个固定长度的特定类型元素组成的序列
 */

 type Currency int

 const (
 	USD Currency = iota		// 0
 	EUR						// 1
 	RMB
 )
func main() {
	// 数组的基本说明
	var a [3]int	// array of 3 integers
	fmt.Println(a[0], a[len(a)-1])
	for i, v := range a {
		fmt.Printf("[%d]%d\n", i, v)
	}
	for _, v := range a {
		fmt.Println(v)
	}

	// 数组声明, 未指定初始值的元素将用零值初始化
	var q [3]int = [3]int {1, 2, 3}
	var r [3]int = [3]int {1, 2}
	fmt.Println(q)
	fmt.Println(r[2])		// default 0

	// []内出现的...，表示数组长度是根据初始化个数来计算的
	p := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("%T\n", p)	// [5]int

	/**
	 [3]int 和 [4]int是两种不同的数组类型
	 数组的长度必须是常量表达式，数组长度需要在编译阶段确定
	 */
	 //p = [3]int{1, 2, 3}		// compile error, 参数类型不统一

	symbol := [...]string{USD:"$", EUR: "€", RMB: "￥"}
	fmt.Println(RMB, symbol[RMB])

	n := [...]int {99: -1, -2}
	fmt.Println(n)

	/*
	 如果数组的元素类型是可以比较的，那么两个数组是可以比较的
	 只有当两个数组的每个元素都相等的时候，两个数组才相等
	 */
	var arr1 = [2]int{1,2}
	var arr2 = [...]int{1,2}
	//var arr3 = [3]int{1,2,3}
	var arr4 = [2]int{1,3}
	fmt.Println(arr1==arr2)	// true
	fmt.Println(arr1==arr4)	// false 第二个元素值不相等
	//fmt.Println(arr4 == arr3)	// compile error, 数组长度不同为不同的数组类型
}
