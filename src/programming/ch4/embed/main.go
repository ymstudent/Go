package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point //匿名成员：不带名称的结构体成员，直接指定类型
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main()  {
	//含有匿名成员的结构体的两种初始化方式
	w := Wheel{Circle{Point{8, 8}, 5}, 20}

	w2 := Wheel{
		Circle: Circle{
			Point: Point{X:8, Y:8},
			Radius: 5,
		},
		Spokes: 20,
	}

	w.X =10
	w2.X = 20
	fmt.Println(w, w2)
}
