package main

import (
	"fmt"
	"math"
)

//使用type声明一个新类型
type MyFloat float64

//我们也可以为这个非结构体类型声明方法
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

//需要注意的是：接收者的类型定义和方法声明必须在同一个包内；不能为内建类型（如float等）声明方法
func main()  {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
