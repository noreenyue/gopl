package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	n := "123456"
	//n = "1234"
	n = "12346788900"
	//n = "+12346788900.1111"
	fmt.Println(comma(n))
}

func comma(s string) string {
	if len(s) <= 0 {
		return s
	}

	// check sign
	sign := ""
	if s[0] == '+' || s[0] == '-' {
		sign = string(s[0])
		s = s[1:]
	}

	// check float
	dot := strings.LastIndex(s, ".")
	floatStr := ""
	if dot >= 0 {
		floatStr = "." + s[dot+1:]
		s = s[:dot]
	}

	n := len(s)
	buf := bytes.Buffer{}
	for i:=0 ; i<n; i++ {
		if (n-i-1) % 3 == 2 && i!=0{
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
	}
	return sign + buf.String() + floatStr
}