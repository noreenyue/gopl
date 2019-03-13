package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(ints2String([] int {1,2,3}))
}

func ints2String(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i >= 1 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}