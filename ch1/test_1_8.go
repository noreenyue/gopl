package main

import (
	"fmt"
	"math/rand"
)

func main() {
	heads, tails := 0, 0
	for i:=0; i<10; i++ {
		switch coinflip() {
		case "heads":
			heads++
		case "tails":
			tails++
		default:
			fmt.Println("landed on edge!")
		}
	}
	fmt.Printf("heads: %d, tails: %d\n", heads, tails)

	// no tag switch = switch true
	x := 0
	for i:=-3; i<7; i++ {
		x += Signum(i)
	}
	fmt.Printf("x=%d\n", x)

	// struct
	var p Point
	p.X = 1
	p.Y = 2
	fmt.Printf("p=%v\n", p)
}

func coinflip() string {
	sides := make(map[int]string)
	sides[0] = "heads"
	sides[1] = "tails"
	return sides[rand.Intn(100)%2]
}

func Signum (x int) int {
	switch {
	case x>0 :
		return +1
	default :
		return 0
	case x<0:
		return -1

	}
}

type Point struct {
	X, Y int
}