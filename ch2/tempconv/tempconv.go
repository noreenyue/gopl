package tempconv

import "fmt"

type Celsius float64	// 摄氏温度
type Fahrenheit float64 // 华氏温度
type Kevin float64 // Kevin温度

const (
	AbsoluteZeroC Celsius = -273.15	// 绝对零度
	FreezingC Celsius = 0			// 结冰温度
	BoilingC Celsius = 100			// 沸腾温度
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (k Kevin) String() string {
	return fmt.Sprintf("%g°K", k)
}