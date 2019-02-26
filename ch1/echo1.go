package main

import (
	"fmt"
	"os"
)


/*
Go语言只有for一种循环方式，for循环有多种方式
  1. for循环的第一种形式，为
	1）标准格式
		for initialization; condition; post {
			// do something
		}

	   2）for循环的三个部分每个都可以省略。如果省略 initialization 和 post，分号也可以省略
		   // a traditional "while" loop
		for condition {
			// ...
		}

	 3) 如果condition也省略，则表示一个无限循环
		// a traditional infinite loop
		for {
			// ...
		}

 2. for循环的第二种形式，range区间 echo2.go
 */

func main() {
	/*
	 * var声明，定义了两个string类型的变量：s和sep
	 * 变量会在声明时直接初始化，如何变量没有显示初始化，会隐式赋值为其类型对应的零值。
	 * 零值：int类型为0，string类型为空字符串""
	 */

	var s, sep string


	// 符号 := 是短变量声明， 定义一个或多个变量，并根据初始值为这些变量赋予适当的类型
	for i := 1; i < len(os.Args); i++ {

		/**
		 * go语言的+=是语句，而不是表达式
		 * 不能用于赋值，即 j=i++ 非法； ++i 非法。
		  */
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
