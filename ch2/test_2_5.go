package main

import (
	"fmt"
	"github.com/noreenyue/gopl/ch2/tempconv0"
)

func main() {
	fmt.Printf("%g\n", tempconv0.BoilingC-tempconv0.FreezingC)
	boilingF := tempconv0.CToF(tempconv0.BoilingC)
	fmt.Printf("%g\n", boilingF-tempconv0.CToF(tempconv0.BoilingC))
	// compile error, type mismatch
	// fmt.Printf("%g\n", boilingF-tempconv0.FreezingC)


	var c tempconv0.Celsius
	var f tempconv0.Fahrenheit
	fmt.Println(c == 0)
	fmt.Println(f >= 0)
	// compile error, type mismatch
	// fmt.Println(f >= c)
	fmt.Println(f >= tempconv0.Fahrenheit(c))

	c2 := tempconv0.FToC(212.0)
	fmt.Println(c2.String())
	fmt.Printf("%v\n", c)	// call String()
	fmt.Printf("%s\n", c)	// call String()
	fmt.Println(c)				// call String()
	fmt.Printf("%g\n", c)	// dose not call String()
	fmt.Println(float64(c))		// dose not call String()
}
