package tempconv

func CToF (c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5+32)
}
func CToK (c Celsius) Kevin {
	return Kevin(c)
}

func FToC (f Fahrenheit) Celsius {
	return Celsius((f-32)/5*9)
}