package main

import (
	"fmt"
	"os"
)



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

	if x:=f2(); x==0 {
		fmt.Println(x)
	} else if y:=g2(); x == y {	// 第一个if中声明的变量，在第二个if中也可以访问
		fmt.Println(x, y)
	} else {
		fmt.Println(x, y)
	}
	// fmt.Println(x, y)	// compile error, x and y are not visible here

	initFunc()
	initFunc2()
}

var cwd string
func initFunc(){
	// 此时赋值的cwd为局部变量, global cwd为unused变量
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stdout,"get cwd failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("current working directory=%v\n", cwd)
}

func initFunc2(){
	var err error
	// cwd 为 global变量
	cwd, err = os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stdout,"get cwd failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("current working directory=%v\n", cwd)
}

/**
	包级别声明的顺序不影响作用域范围
 */
func f2() int{
	return 0
}

func g2() int{
	return 0
}