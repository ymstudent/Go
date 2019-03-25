package main

import (
	"fmt"
	"math"
)

//接口值
//接口也是值，也可以像其他值一样传递。接口值可以用作函数的参数或返回值
//在内部，接口值可以看作包含值和具体类型的元组：（value, type）
//接口值保存了一个具体底层类型的具体值，接口值调用方法时会执行其底层类型的同名方法

type I interface { //这是接口
	M()
}

type T struct {
	S string
}

func (t *T) M()  {
	fmt.Println(t.S)
}

type F float64

func (f F) M()  {
	fmt.Println(f)
}

func main()  {
	var i I
	i = &T{"Hello World!"} //这是接口值
	describe(i)
	i.M() //接口值i执行了类型*T的M()方法

	i = F(math.Pi)
	describe(i)
	i.M() //执行F类型的M方法

	//个人是这样理解接口值的：实现了接口的类型T，类型T的具体值被赋给变量i，这个变量i就叫接口值，它的底层是一个元组
}

func describe(i I)  {
	fmt.Printf("(%v, %T)\n", i, i)
}
