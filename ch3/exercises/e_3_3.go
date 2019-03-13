package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "heello"
	s2 := "heallo"
	s1 = "loleh"
	s2 = "ellho"
	fmt.Println(isDisOrder(s1, s2))
}

func isDisOrder(s1 string, s2 string) bool {
	if len(s1) != len(s2){
		return false
	}

	for _, c := range s1{
		idx := strings.LastIndex(s2, string(c))
		if idx < 0 {
			return false
		}
		s2 = s2[0:idx] + s2[idx+1:]	// del
	}
	return true
}