package main

import "fmt"

//Stringer
//fmt包中定义的Stringer是最普遍的接口之一
//	type Stringer interface {
//    	String() string
//	}
//Stringer是一个可以用字符串描述自己的类型。fmt包通过此接口来打印值

type Vertex struct {
	Name string
	Age int
}

//类型*Vertex实现了Stringer接口
func (v *Vertex) String() string {
	return fmt.Sprintf("%v (%v years)", v.Name, v.Age)
}

func main() {
	a := &Vertex{"Tom", 18}
	b := &Vertex{"Bob", 24}
	fmt.Println(a, b)
}
