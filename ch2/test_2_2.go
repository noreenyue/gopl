package main

import "fmt"

func main() {
	i := 1
	j := 2
	fmt.Printf("before swap: i=%d, j=%d\n",i ,j )
	j, i = i, j // swap
	fmt.Printf("after swap: i=%d, j=%d\n",i ,j )

	x := 1
	p := &x
	*p = 2
	fmt.Println(x)

	var m, n int
	// %t：bool类型
	fmt.Printf("&m==&m:%t, &m==&n:%t, &m==nil:%t\n", &m==&m, &m==&n, &m==nil )

	var p1 = f(m)
	fmt.Printf("point at=%d\n", p1)

	v1 := 1
	fmt.Printf("incr v1 from %d to ", v1)
	incr(&v1)
	fmt.Printf("%d\n", v1)
}

func f(v int) *int {
	return &v
}


func incr(p *int) int{
	*p ++
	return *p
}
