package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

/**
 字符串
 */

func main() {
	s := "hello world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])

	//fmt.Println(s[len(s)])	// panic 越界

	// 切片
	fmt.Println(s[0:5])
	fmt.Println(s[:5])
	fmt.Println(s[6:])
	fmt.Println(s[:])
	fmt.Println("goodbye, " + s[6:])	// 拼接字符串

	s1 := "left foot"
	t := s1
	s1 += ", right foot"
	fmt.Println(t)
	fmt.Println(s1)

	//s[0] = 'L' // compile error 字符串不可被修改，不能被赋值

	s = `\n`	// 原生字符串，转义无效
	fmt.Println(s)

	us := "hello, 世界"
	fmt.Println(len(us))	// 13 7+3*2 一个unicode占3个字节
	fmt.Println(us[7:])
	fmt.Println(us[8:])	// 乱码
	fmt.Println(us[10:])

	for i := 0 ; i < len(us) ; {
		r, size := utf8.DecodeRuneInString(us[i:])
		fmt.Printf("%c[%d],", r, i)
		i += size
	}
	fmt.Println()

	// range 会自动隐式处理unicode字符
	for i, r := range us {
		fmt.Printf("%c[%d],", r, i)
	}
	fmt.Println()

	count := 0
	for _, _ = range us {
		count ++
	}
	fmt.Println(count, utf8.RuneCountInString(us), count==utf8.RuneCountInString(us))

	us2 := "我是歪歪"
	fmt.Printf("% x\n", us2)
	// string接收转换至rune类型
	r := []rune(us2)
	fmt.Printf("%x\n", r)
	fmt.Println(string(r))

	fmt.Println(string(65))
	fmt.Println(string(0x5cb3))	// 岳


	str := "abc"
	b := []byte(str)
	s2 := string(b)
	fmt.Println(b, s2)

	// 3.5.5
	x := 124
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))
	// 指定进制来格式化数字
	fmt.Println(strconv.FormatInt(int64(x), 2))
	fmt.Println(strconv.FormatUint(uint64(x), 2))

	n1, err1 := strconv.Atoi("123")
	n2, err2 := strconv.ParseInt("123", 10 ,64)
	fmt.Println(n1, n2, err1, err2)
}

func startsWith (s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func endsWith (s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(suffix):] == suffix
}