package main

import "fmt"

//接口与隐式实现
//类型通过实现一个接口所有的方法来实现该接口。既然无需专门的显示声明，也就没有“implements”关键字
//隐式接口从接口的实现中解耦了定义，这样接口的实现可以出现在任何包中，无需提前准备
//因此，也就无需再每一个实现上增加新的接口名称，这样同时也鼓励了明确的接口定义

type I interface {
	M()
}

type T struct {
	S string
}
//此方法表示类型T实现了接口I，但是我们无需显示的声明（T implements I）
func (t T) M()  {
	fmt.Println(t.S)
}
func mian()  {
	var i I = T{"Hello World!"}
	i.M()
}
