package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func cost1() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func cost2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:]{
		s += arg + sep
		sep = " "
	}
	fmt.Println(s)
}


func main() {
	t1 := time.Now()
	for i:=0 ; i<100; i++ {
		cost1()
	}
	t2 := time.Now()

	t3 := time.Now()
	for i:=0 ; i<100; i++ {
		cost2()
	}
	t4 := time.Now()

	// func cost2 cost more
	fmt.Println( t2.Sub(t1))
	fmt.Println( t4.Sub(t3))
}