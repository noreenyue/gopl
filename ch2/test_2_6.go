package main

import (
	"fmt"
	"github.com/noreenyue/gopl/ch2/tempconv"
)

func main() {
	fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC)
	fmt.Println(tempconv.CToF(tempconv.BoilingC))

	fmt.Println(tempconv.CToK(tempconv.AbsoluteZeroC))
}
