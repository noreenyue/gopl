package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/*
	运行：go run ch1/dup3.go ch1/contents.txt
	不同于dup1和dup2，dup3读取整个文件内容，再逐行打印
 */

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3 %v\n", err)
		}

		// ReadFile返回类型是字节切片(byte splice)，需强转为string类型
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for txt, num := range counts {
		if num>1 {
			fmt.Println(txt, ":", num)
		}
	}
}