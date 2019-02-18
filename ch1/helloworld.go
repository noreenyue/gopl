package main
// main包定义一个独立可执行程序，而不是一个库

// GO语言提供的标准库，用于格式化输入输出
import "fmt"

/**
1. 可以直接执行文件
    go run helloworld.go

2. 也可以先编译再执行
    1) go build helloworld.go 生产一个helloworld的二进制文件
    2) 直接运行可执行文件 ./helloworld
*/

// main包里的main函数，是整个程序执行的入口
func main() {
    // GO语言原生支持Unicode
    fmt.Println("Hello, 世界！")
}
