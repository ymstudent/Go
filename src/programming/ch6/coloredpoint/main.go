package main

import "image/color"

type Point struct {
	X, Y float64
}

type ColoredPoint struct {
	Point
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
}
