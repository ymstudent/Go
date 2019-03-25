package main

import "fmt"

func main()  {
	//类型 [n]T 表示拥有 n 个 T 类型的值的数组。在GO中，数组的长度是不可更改的
	a := [3]int{1,2,3}
	fmt.Println(a)
	fmt.Printf("type: %T\n", a)
	b := [3]int{}
	fmt.Println(b)
	fmt.Printf("type: %T\n", b)
}
