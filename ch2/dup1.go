package main

import (
	"bufio"
	"fmt"
	"os"
)


/**
 	运行：go run ch2/dup1.go < ch2/contents.txt
	说明：每次执行，打印出来的顺序不固定，map无序
 */
func main() {
	// 定义一个空的map，key为string类型，value为int类型
	items := make(map[string] int)

	// 控制台输入，没有break条件，会一直循环
	// 文件输入，流读不到后会中断
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		items[input.Text()] ++
	}

	// items是map类型，随机顺序打印
	for text, num := range items {
		print(fmt.Sprintf("%s\t%d\n", text, num))
	}

}

