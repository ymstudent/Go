package main

//练习：定义你自己的 Image 类型，实现必要的方法并调用 pic.ShowImage。

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {}

/**颜色模式*/
func (img Image) ColorModel() color.Model  {
	return color.RGBAModel
}

/**图片边界*/
func (img Image) Bounds() image.Rectangle  {
	return image.Rect(0,0,200,200)
}

/**某个点的颜色*/
func (img Image) At(x, y int) color.Color  {
	return color.RGBA{uint8(x),uint8(y),255,255}
}

func main()  {
	m := Image{}
	pic.ShowImage(m)
}
