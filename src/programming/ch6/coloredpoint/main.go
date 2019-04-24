package main

import (
	"fmt"
	"image/color"
)

type Point struct {
	X, Y float64
}

type ColoredPoint struct {
	Point
	color color.RGBA
}

type ColoredPoint2 struct {
	*Point
	color color.RGBA
}

func (p *Point) ScaleBy(factor float64)  {
	p.X *= factor
	p.Y *= factor
}

func main()  {
	//可以直接调用内嵌Point类型的方法，而不需要提到Point类型，因为Point类型的方法都被纳入到ColoredPoint类型中
	red := color.RGBA{255,0,0,255}
	p := ColoredPoint{Point{1, 2}, red}
	p.ScaleBy(2)
	p.Point.ScaleBy(2)
	//通过指针类型的匿名字段，共享通用的结构使对象之间的关系更加动态，多样化
	bule := color.RGBA{0,0,255,255}
	a := ColoredPoint2{&Point{1,1}, bule}
	b := ColoredPoint2{&Point{5,4}, bule}
	a.Point = b.Point //a, b共享同一个Point
	p.ScaleBy(2)
	fmt.Println(a.Point, b.Point) //"{2,2} {2,2}"
	//编译器处理内嵌结构体选择子的优先级：
	//当前结构体选择子（eg:ColoredPoint的选择子） > 内嵌字段的选择子（eg:ColoredPoint.Point.X的选择子） > 内嵌字段的字段的选择子（eg:ColoredPoint.Point.X的选择子）
}
