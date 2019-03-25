package main

import "fmt"

func main()  {
	//Go 拥有指针。指针保存了值的内存地址。
	//类型 *T 是指向 T 类型值的指针。其零值为 nil。
	//& 操作符会生成一个指向其操作数的指针。
	a := 1
	p := &a
	fmt.Printf("type: %T\n", p)
	fmt.Printf("value: %p\n", p)
	//* 操作符表示指针指向的底层值。这也就是通常所说的“间接引用”或“重定向”。
	fmt.Println(*p)
	*p = 2
	fmt.Println(a)
}