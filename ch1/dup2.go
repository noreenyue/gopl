package main

import (
	"bufio"
	"fmt"
	"os"
)
/*
	运行：go run ch1/dup2.go ch1/contents.txt
	dup1 和 dup2 以流形式读取文件
 */
func main() {
	lines := make(map[string]int)

	files := os.Args[1:]
	if len(files) == 0 {
		countFileLines(os.Stdin, lines)
	} else {
		for _, fileName := range files {
			// 打开文件，变量类型 (*os.File)
			f, err := os.Open(fileName)

			// 打开失败 nil表示空
			if err != nil {
				fmt.Fprintf(os.Stderr, "open file error: %v\n", err)
			}
			countFileLines(f, lines)
			// 关闭文件
			f.Close()
		}
	}

	for txt, num := range lines {
		if num > 1 {
			fmt.Println(txt, ": ", num)
		}
	}
}

func countFileLines(f * os.File, lines map[string] int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		lines[input.Text()]++
	}
}