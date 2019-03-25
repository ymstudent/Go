package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

//现在我们将Abs和Scale方法重写为函数
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64)  {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Scale2(f float64)  {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}


func main()  {
	v := Vertex{3, 4}
	//如果我们将此处的&去掉，会发生编译错误。
	//通过比较Scale函数和Scale2方法，可以发现：带指针参数的函数必须接受一个指针，而以指针为接收者的方法被调用时，接收者既能为指针也能为值
	Scale(&v, 10)
	fmt.Println(Abs(v))
	//对于语句v.Scale2(10)，即便 v 是个值而非指针，带指针接收者的方法也能被直接调用。
	//也就是说，由于 Scale2 方法有一个指针接收者，为方便起见，Go 会将语句 v.Scale(5) 解释为 (&v).Scale(5)。
	v.Scale2(10)
	fmt.Println(Abs(v))

	//同样的事情也发生在相反的方向。
	//接受一个值作为参数的函数必须接受一个指定类型的值。
	v = Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))
	//而以值为接收者的方法被调用时，接收者既能为值又能为指针。这种情况下，方法调用 p.Abs() 会被解释为 (*p).Abs()。
	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))

	//选择使用指针作为接收者的原因有2个：
	//1是方法能够修改其接收者指向的值
	//2是这样可以避免每次调用方法时复刻该值。如果值的类型为大型结构体时，这样做会更加高效
	//需要注意的是：所有给定类型的方法都应该有值或者指针接收者。但是不因该将二者混用
}
