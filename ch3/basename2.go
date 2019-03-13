package main

import (
	"fmt"
	"strings"
)

func main() {
	path := "/home/xiaoju/yueyuan.log"
	path = "yueyuan.log"
	path = "/home/xiaoju/"
	path = ""
	path = "y"
	path = "/"
	fmt.Println(basename2(path))
}

func basename2(path string ) string {
	i := strings.LastIndex(path, "/")
	if dot := strings.LastIndex(path, "."); dot > 0{
		path = path[:i]
	}
	return path
}
