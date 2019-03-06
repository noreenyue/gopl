package main
/**
	go run ch2/echo4.go 1 2 3
	go run ch2/echo4.go -s=/ 1 2 3
	go run ch2/echo4.go -s=/ -n=false 1 2 3
	go run ch2/echo4.go -s=/ -n=true 1 2 3
	go run ch2/echo4.go -s=/ -n=1 1 2 3

	// exception examples
	go run ch2/echo4.go -s=/ -n=e 1 2 3
	go run ch2/echo4.go -help
 */
import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Printf("%s\n", strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
		fmt.Printf("n=%v\n", *n)
	}
}
