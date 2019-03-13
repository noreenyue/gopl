package main

import "fmt"

func main() {
	path := "/home/xiaoju/yueyuan.log"
	path = "yueyuan.log"
	path = "/home/xiaoju/"
	path = ""
	path = "y"
	path = "/"
	fmt.Println(basename(path))
}

func basename(path string) string {
	// 先截断路径
	for i := len(path) - 1 ; i >= 0; i-- {
		if path[i] == '/' {
			path = path[i+1:]
			break
		}
	}

	for i := len(path) - 1 ; i >= 0; i-- {
		if path[i] == '.' {
			path = path[:i]
			break
		}
	}
	return path
}
