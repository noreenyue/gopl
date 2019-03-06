package main

import "fmt"

var globalP *int

func main() {
	p := new(int)
	fmt.Println(*p)
	*p = 2
	fmt.Printf("%d\n", *p)

	p2 := newInt()
	fmt.Println( *p2)

	p3 := newInt2()
	fmt.Println( *p3)

	// new只在函数内部被重新定义
	p4 := new(int)
	fmt.Println(delta(*p4, 5))
}

func newInt() *int {
	return new(int)
}

func newInt2() *int {
	v := 1
	// 局部变量作用域被扩大到global，叫做变量逃逸
	// 函数返回后，局部变量依然可达，必须在堆上分配该变量
	globalP = &v
	return &v
}

// new被定义为其他类型
func delta(old int, new int) int {
	// exception: cannot call non-function new
	// fmt.Println(new(int))
	return new - old
}

/**
	虽然是new出来的变量，作用域只在函数内.
	g()函数返回后，变量y立即被回收。
	编译器直接在栈上分配该变量即可，也可以选择在堆上分配，然后由GC回收变量内存空间
 */
func g() {
	y := new(int)
	*y = 1
	fmt.Println(*y)
}