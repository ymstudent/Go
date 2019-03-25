//surface函数根据一个三维曲面函数计算并生成SVG
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

var width, height float64 = 600, 320      //以像素表示的画布大小
var cells float64 = 100                   //网格单元的个数
var xyrange float64 = 30.0                //坐标轴的范围(-xyrange...+xyrange)
var xyscale float64 = width / 2 / xyrange //x或y在每个单位长度的像素
var zscale float64 = height * 0.4         //z轴上每个单位长度的像素
var angle float64 = math.Pi / 6           //x, y轴的角度 (=30°)

/*const (
	width, height = 600, 320            //以像素表示的画布大小
	cells         = 100                 //网格单元的个数
	xyrange       = 30.0                //坐标轴的范围(-xyrange...+xyrange)
	xyscale       = width / 2 / xyrange //x或y在每个单位长度的像素
	zscale        = height * 0.4        //z轴上每个单位长度的像素
	angle         = math.Pi / 6         //x, y轴的角度 (=30°)
)*/

var sin30, cos30 = math.Sin(angle), math.Cos(angle) //sin(30°), cos(30°)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil)) //练习3.4
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	if err := r.ParseForm(); err != nil {
		return
	}

	for k, v := range r.Form {
		if k == "height" {
			h, _ := strconv.ParseFloat(v[0], 64)
			if h > 0 {
				height = h
			}
		}
		if k == "width" {
			w, _ := strconv.ParseFloat(v[0], 64)
			if w > 0 {
				width = w
			}
		}
	}

	xyscale = width / 2 / xyrange
	zscale = height * 0.4

	//定义线条为红色，单元格为蓝色。当在峰顶时，单元格小，线条密集，呈现红色，谷底单元格大，呈现蓝色：练习3.3
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: #ff0000; fill: #0000ff; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)

	for i := 0; i < int(cells); i++ {
		for j := 0; j < int(cells); j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Fprintf(w, "<polygon points='%g, %g, %g, %g, %g, %g, %g, %g' />\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}

	fmt.Fprintf(w, "</svg>")
}

func corner(i, j int) (float64, float64) {
	//求出网格单元(i, j)的顶点坐标 (x, y)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	//计算曲面高度z
	z := f(x, y)

	//将(x, y, z)等角投射到二维SVG绘图平面上，坐标是(sx, sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) //到(0,0)的距离
	//r := math.Dim(x, y) //练习3.2
	//r := math.Atan2(x, y)
	//r := math.Max(x, y)
	//r := math.Copysign(x, y)
	//r := math.Min(x, y)
	//r := math.Remainder(x, y)
	if math.IsInf(math.Sin(r)/r, 0) {
		return math.MaxFloat64 //如果计算结果是无穷大，则返回float的最大值：练习3.1
	}
	return math.Sin(r) / r
}
