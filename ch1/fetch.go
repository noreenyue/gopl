package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

/*
	运行： go run ch1/fetch.go http://www.baidu.com	// 读取网站的内容
		  go run ch1/fetch.go www.baidu.com		// http error，无法识别协议
 */

func main() {
	for _, url := range os.Args[1:] {
		// 创建http请求resp, resp的body字段包括一个可读的服务器响应流
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stdout, "http get error: %v\n", err)
			os.Exit(1)
		}

		// 将resp.Body的内容全部读取，结果保存到byte变量b中
		b, err := ioutil.ReadAll(resp.Body)
		// 关闭流，防止资源泄露
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stdout, "read from response body error: %v\n", err)
			os.Exit(1)
		}
		// 将b写入标准输出流
		fmt.Printf("%s", b)
	}
}
