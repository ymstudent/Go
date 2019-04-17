//匿名函数
package main

import "fmt"

func main()  {
	f := squares()
	fmt.Println(f()) //1
	fmt.Println(f()) //4
	fmt.Println(f()) //9
	fmt.Println(f()) //16
}

func squares() func() int {
	var x int
	return func() int {
		fmt.Printf("squares x is: %v\n", x)
		x++
		return x*x
	}
}
