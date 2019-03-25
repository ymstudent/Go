package main

import (
	"fmt"
)

//Go函数可以是一个闭包。闭包是一个函数值，它引用了其函数体之外的变量。该函数可以访问并赋予其引用的变量的值，换句话说，该函数被“绑定”在了这些变量上
func adder() func(int) int {
	sum := 0
	return func(x int) int { //引用了函数体之外的的变量sum
		sum += x
		return sum
	}
}

func main()  {
	//函数adder返回了一个闭包。每个闭包都被绑定在其各自的sum变量上
	pos, neg := adder(), adder()
	for i:= 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}
