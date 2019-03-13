package main

import (
	"fmt"
	"math"
	"time"
)

const noDelay time.Duration = 0
const timeout = time.Minute * 5

const (
	ga = 1
	gb	// 默认使用前一个初始化表达式
	gc = 2
	gd
)


type Weekday int
const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type Flags int
const (
	FlagUp Flags = 1 << iota
	FlagBroadcast // supports broadcast access capability
	FlagLoopback // is a loopback interface
	FlagPointToPoint // belongs to a point-to-point link
	FlagMulticast // supports multicast access capability
)

const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB
	ZiB
	YiB
)
func main() {
	fmt.Printf("%T %[1]v\n", noDelay)
	fmt.Printf("%T %[1]v\n", timeout)
	fmt.Printf("%T %[1]v\n", time.Minute)

	fmt.Println(ga, gb, gc, gd)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
	fmt.Println(FlagUp, FlagBroadcast, FlagLoopback, FlagMulticast, FlagPointToPoint)

	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v))
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))
	SetBroadCast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))
	fmt.Printf("%b %t\n", v, isCast(v))

	fmt.Println(YiB/ZiB)

	// math.Pi是一个无类型浮点数常量
	const Pi64 float64 = math.Pi
	var f32 float32 = float32(Pi64)
	var f64 float64 = Pi64
	var fcmpl complex128 = complex128(Pi64)
	fmt.Println(f32, f64, fcmpl)

	var f float32 = 212
	fmt.Println((f-32)*5/9) // int型
	fmt.Println(5/9*(f-32))	// 0
	fmt.Println(5.0/9.0*(f-32))	// float

	// 无类型复数隐式转换
	var f2 float64 = 3 + 0i
	f2 = 2
	fmt.Println(f2)
	f2 = 1e123
	fmt.Println(f2)
	f2 = 'a'
	fmt.Println(f2)

	fmt.Printf("%T %[1]v\n", 0)
	fmt.Printf("%T %[1]v\n", 0.0)
	fmt.Printf("%T %[1]v\n", 0x000)
	fmt.Printf("%T %[1]v\n", 0i)
	fmt.Printf("%T %[1]v\n", '\000')
}


func IsUp(v Flags) bool {
	return v&FlagUp == FlagUp
}

func TurnDown(v *Flags) {
	*v &^= FlagUp
}

func SetBroadCast(v *Flags) {
	*v |= FlagBroadcast
}

func isCast(v Flags) bool {
	return v & (FlagBroadcast | FlagMulticast) != 0
}