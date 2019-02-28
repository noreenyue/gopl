package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
	运行：go run ch1/exercises/e_1_4.go  ch1/contents.txt
*/

func main() {
	counts := make(map[string]int)
	filenames := os.Args[1:]
	if len(filenames) == 0 {
		fmt.Fprintf(os.Stderr, "no file found to open: %v \n")
	} else {
		for _, filename := range filenames {
			f, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "file open failed: %v \n", err)
			}
			countFileLines(f, counts)

			if len(counts) > 1 {
				fmt.Println(filename)
			}
		}
	}
}

func countFileLines (f * os.File, counts map[string] int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()] ++
	}
}