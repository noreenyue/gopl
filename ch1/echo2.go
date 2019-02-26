package main

import (
	"fmt"
	"os"
)
/*
	Go语言for循环的第二种形式，range区间
	每次循环迭代时，range产生一对索引和元素值。
		for _, arg := range argList {
			// ...
		}
 */

 /*
 	变量的几种声明方式，前两种较常见
 	1) s := ""
 		短变量声明，最简洁。只允许局部变量，不可以用做包变量
 	2) var s string
 		隐式赋值为string类型的零值，即空字符串""
 	3) var s = ""
 		很少用，除非用于多个变量声明
 	4) var s string = ""
 		冗余
  */

func main() {
	s, sep := "", " "
	// _ 空标识符，用于语法需要而程序逻辑不需要
	// Go语言语法不允许不被使用的局部变量，会导致编译错误
	for _, arg := range os.Args {
		// 每次循环迭代，s的内容会被更新。原来的内容会被垃圾回收，如果连接涉及的数据量较大，代价较高
		s += sep + arg
	}
	fmt.Println(s)
}