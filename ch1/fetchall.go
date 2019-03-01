package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

// main方法本身也运行在一个goroutine里
func main() {
	start := time.Now()
	// channel类型是goroutine之间用来传递参数的
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		// go 关键字用来创建一个goroutine，goroutine是go里一种函数的并发执行方式
		go fetchContents(url, ch)
	}

	for i:=0; i<len(os.Args[1:]); i++{
		// <- ch  , receive from channel
		fmt.Println(<-ch)
	}

	cost := time.Since(start).Seconds()
	fmt.Printf("main cost time: %0.2fs \n", cost)
}

func fetchContents(url string, ch chan<- string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		// ch <-  , send to channel
		ch <- fmt.Sprintf("http get error: url=%s, %v\n", url, err)
		return
	}

	// ioutil.Discard 可看做是 /dev/null，用来接收一些不需要保存的数据
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("resp body read error: url=%s, %v\n", url, err)
		return
	}
	resp.Body.Close()

	cost := time.Since(start).Seconds()
	ch <- fmt.Sprintf("fetch url cost time: %0.2fs, read bytes,: %d, url:%s ", cost, nbytes, url)
}
