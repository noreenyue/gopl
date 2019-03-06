package main

import (
	"fmt"
	"github.com/noreenyue/gopl/ch2/tempconv"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stdout, "arg parse float error: %v\n", err)
		}

		c := tempconv.Celsius(t)
		f := tempconv.Fahrenheit(t)

		// 调用了String()函数
		fmt.Printf("%s=%s, %s=%s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))

	}
}
