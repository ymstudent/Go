package main

import (
	"fmt"
	"math"
)

//Go中没有类。到那时可以为结构体类型定义方法
//方法是一类带有特殊的接收者参数的函数
//方法接收者在它自己的参数列表内，位于func关键字和方法名之间
type Vertex struct {
	X, Y float64
}
//下面这个Abs方法拥有一个名为v，类型为Vertex的接收者
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//记住：方法只是个带接收者参数的函数。现在我们写一个Abs2函数，它和上面的方法相比，功能上并没有什么变化
func Abs2(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main()  {
	v := Vertex{3,4}
	fmt.Println(v.Abs())
	fmt.Println(Abs2(v))
}
