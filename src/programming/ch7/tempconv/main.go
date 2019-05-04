package main

import (
	"flag"
	"fmt"
	"programming/ch2/tempconv"
)

//*celsiusFlag满足flag.Value接口
type celsiusFlag struct {
	tempconv.Celsius
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s,"%f%s", &value, &unit)
	switch unit {
	case "C", "℃":
		f.Celsius = tempconv.Celsius(value)
		return nil
	case "F", "℉":
		f.Celsius = tempconv.FToC(tempconv.Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

//CelsiusFlag 根据给定的name、默认值和使用方法定义了一个Celsius标志，返回了标志值的指针
func CelsiusFlag(name string, value tempconv.Celsius, usage string) *tempconv.Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main()  {
	flag.Parse()
	fmt.Println(*temp)
}
