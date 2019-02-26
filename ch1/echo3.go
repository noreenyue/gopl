package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// 高效的字符串拼接方式
	fmt.Println(strings.Join(os.Args[1:], " "))
}