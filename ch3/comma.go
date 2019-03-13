package main

import "fmt"

func main() {
	n := "123456"
	n = "1234"
	n = "12346788900"
	fmt.Println(comma(n))
}


func comma(s string) string {
	if len(s) <= 3 {
		return s
	}
	return comma(s[:len(s)-3])+","+s[len(s)-3:]
}