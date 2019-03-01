package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		// 判断url前缀
		if !strings.HasPrefix(url,"http://") {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stdout, "http get error: %v\n", err)
			os.Exit(1)
		}

		// 直接resp.Body中的内容copy到标准输出流
		len, err := io.Copy(os.Stdout, resp.Body);
		if err != nil {
			fmt.Fprintf(os.Stdout, "response body readall error: %v\n", err)
			os.Exit(1)
		}
		// resp.Status, http 协议状态码
		fmt.Printf("http status: %s, read %d bytes\n", resp.Status, len)
	}
	
}
