package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex;

var count int;
func main() {
	// 除了count请求，其他都走到默认
	http.HandleFunc("/", echoPath)
	// count请求走conter方法
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}


func echoPath(w http.ResponseWriter, r * http.Request) {
	// 需限制每次修改计数的，只有一个goroutine
	mu.Lock()
	count ++
	mu.Unlock()

	fmt.Fprintf(w, "%s, %s, %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
	}

	fmt.Fprintf(w,"Host=%q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr=%q\n", r.RemoteAddr )

	if err := r.ParseForm() ; err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q]=%q\n", k, v)
	}
}

func counter (w http.ResponseWriter, r * http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "count=%d\n", count)
	mu.Unlock()
}