package main

import "fmt"

//func f() {}

var gl = "gl"
func main() {
	f := "f"
	// 局部变量 f 屏蔽了包级别函数f()
	fmt.Printf("%v\n", f)
	fmt.Printf("%v\n", gl)
	//fmt.Println("%v\n", h)	// compile error

	// 每一次声明的变量v，作用域不一样
	v := "hello!"
	for i:=0; i<len(v); i++ {
		v := v[i]
		if v != '!' {
			// 赋值左边的变量v的初始化操作，引用了外部作用域的变量v
			v := v + 'A' - 'a'	// 转大写
			fmt.Printf("%c", v)
		}
	}
	fmt.Println()

	u := "hello!"
	// 隐式声明变量u
	for _, u := range u {
		u := u + 'A' - 'a'
		fmt.Printf("%c", u)
	}
	fmt.Println()


}
