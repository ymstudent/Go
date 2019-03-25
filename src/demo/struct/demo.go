package main

import "fmt"

type Vertex struct{
	X int
	Y int
}

var (
	v1 = Vertex{1,2} //创建一个 Vertex 类型的结构体
	v2 = Vertex{X:1} //Y:0 被隐式地赋予
	v3 = Vertex{} //X:0 Y:0
	v4 = &Vertex{1,2} //创建一个 *Vertex 类型的结构体（指针）
)

func main()  {
	v := Vertex{1,2}
	//结构体字段使用点号来访问。
	v.X = 4
	fmt.Println(v)
	//结构体字段可以通过结构体指针来访问。
	p := &v
	p.X = 10
	fmt.Printf("type: %T\n", p)
	fmt.Printf("value: %p\n", p)
	fmt.Println(v)
	fmt.Println(v1, v2, v3, v4)
}