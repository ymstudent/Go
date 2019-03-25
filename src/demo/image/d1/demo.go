package main

import (
	"fmt"
	"image"
)

//image包定义了Image接口
//	package image
//
//	type Image interface {
//    	ColorModel() color.Model
//    	Bounds() Rectangle
//    	At(x, y int) color.Color
//	}

func main()  {
	m := image.NewNRGBA(image.Rect(0,0,100,100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0,0).RGBA())
}
