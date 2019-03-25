package main

import "fmt"

//nil接口值
//nil接口值既不保存值也不保存具体类型
//为nil接口调用方法时会产生运行时错误，因为接口内的元组并未包含能够指明该调用那个具体方法的类型

type I interface {
	M()
}

func main()  {
	var i I
	describe(i)
	i.M() //panic: runtime error: invalid memory address or nil pointer dereference
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}