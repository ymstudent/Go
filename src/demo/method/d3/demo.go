package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
//为*Vertex指针声明一个方法。指针接受者的方法介意修改接收者指向的值。对于需要经常修改其接收者的方法，指针接收者比值接收者更常用
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
//对比
func (v *Vertex) Scale2(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
func main()  {
	v := Vertex{3,4}
	//若使用值接收者，那么 Scale2 方法会对原始 Vertex 值的副本进行操作。（对于函数的其它参数也是如此。）
	v.Scale2(10)
	fmt.Println(v)
	fmt.Println(v.Abs()) //输出5，此处Vertex的值没有任何改变
	//Scale 方法必须用指针接受者来更改 main 函数中声明的 Vertex 的值。
	v.Scale(10)
	fmt.Println(v)
	fmt.Println(v.Abs()) //输出50，此处Vertex的值出现了改变
}
