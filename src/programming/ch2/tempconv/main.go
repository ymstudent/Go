package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

func (c Celsius) String() string  {
	return fmt.Sprintf("%g℃", c)
}
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g℉", f)
}

//摄氏温度转为华式温度
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

//华氏温度转为摄氏温度
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}