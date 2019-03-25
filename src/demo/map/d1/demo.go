package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

//映射将键映射到值。映射的零值为 nil 。nil 映射既没有键，也不能添加键。
var m map[string]Vertex

func main()  {
	//make 函数会返回给定类型的映射，并将其初始化备用。
	fmt.Printf("type: %T\n", m)
	fmt.Println(m)
	if m == nil {
		fmt.Println("nil!") //此处会输出
	}
	m = make(map[string]Vertex) //此处不可注释，但是上面的m声明却可以，为什么？因为上面申明的是一个nil映射，不能添加键
	fmt.Printf("type: %T\n", m)
	fmt.Println(m)
	if m == nil {
		fmt.Println("nil!") //此处不会输出
	}
	m["Bell Labs"] = Vertex{40.68433, -74.39967}
	fmt.Println(m["Bell Labs"])
}
