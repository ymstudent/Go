//方法：某种特定类型的函数（简单的说就是将函数绑定到特定的参数的类型上）
package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

//普通包级别的函数，可以看做main.Distance
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//类型Point的方法，由于每个类型都有自己的命名空间，索引这里可以看做p.Distance
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

/*
//由于方法与字段来自于同一个命名空间（此处为Point），所以如果方法和字段重名，编译器会报错
func (p Point) X(q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}
*/

/*
//同一包下任何类型都可以声明方法，除了指针类型与接口类型
type inter interface {}
type Pointer *Point
func (i inter) Test(s string) string {
	return s
}
func (p Pointer) Test(s string) string {
	return s
}
*/

func main() {
	//计算周长
	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())
}
