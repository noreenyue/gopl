package tempconv0

import "fmt"

/*
	名字是大写的，则该名字是导出的，可以在外部被调用
	变量、方法、常量
*/
type Celsius float64	// 摄氏温度
type Fahrenheit float64 // 华氏温度

const (
	AbsoluteZeroC Celsius = -273.15	// 绝对零度
	FreezingC Celsius = 0			// 结冰温度
	BoilingC Celsius = 100			// 沸腾温度
)

func CToF (c Celsius) Fahrenheit{
	return Fahrenheit(c*9/5+32)
}

func FToC (f Fahrenheit) Celsius {
	return (Celsius(f-32))*5/9
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}