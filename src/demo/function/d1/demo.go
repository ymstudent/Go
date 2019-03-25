package main

import (
	"fmt"
	"math"
)

//函数也是值，也可以像其他值一样传递，函数值可以用作函数的参数或返回值
func computer(fn func(float64, float64) float64) float64 {
	return fn(3,4)
}

func main()  {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y) //开平方计算
	}
	fmt.Println(hypot(5,12))

	fmt.Println(computer(hypot)) //此处相当于hypot(3,4)，将hypot函数作为参数传入了computer
	fmt.Println(computer(math.Pow)) //math.Pow(x,y)返回值为y的x次方
}
