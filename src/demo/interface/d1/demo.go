package main

import (
	"fmt"
	"math"
)

//接口类型是由一组方法签名定义的集合，接口类型的变量可以保存任何实现了这些方法的值
type Abser interface {
	Abs() float64
}

func main()  {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3,4}

	a = f // a MyFloat 实现了Abser
	a = &v // a *Vertex 实现了Abser
	//此处由于v 是一个 Vertx(而不是*Vertex),而Vertex没有实现Abser。因此编译是会报错
	//此时，如果过我们将41行的*去掉就能编译通过了，因为以值为接收者的方法被调用时，接收者既能为指针，也能为值
	a = v

	fmt.Println(a.Abs())
}


type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
