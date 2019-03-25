package main

import "fmt"

//类型选择是一种按顺序从几个类型断言中选择分支的结构
//类型选择与switch语句相似，不过类型选择中的case为类型（而非值），它们针对给定接口值所存储的值的类型进行比较

func do(i interface{})  {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main()  {
	do(21)
	do("hello")
	do(true)
}
