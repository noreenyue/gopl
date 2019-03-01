package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
	后台启动 server go run ch1/server1.go &
	浏览器执行 http://127.0.0.1:8000/hello
 */
func main() {
	// 将发送到/的请求都调用handler方法
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler( w http.ResponseWriter, r * http.Request) {
	fmt.Fprintf(w, "Url.Path=%q\n", r.URL.Path)
}
