package main

import "fmt"
//练习：实现斐波那契闭包（f(n) = f(n-1) + f(n-2)）
//返回一个"返回int的闭包"
func fibonacci() func(int) int {
	f0 := 0
	f1 := 1
	return func(x int) int {
		if x < 2 {
			return x
		}
		v := f0 + f1
		f0 = f1
		f1 = v
		return v
	}
}

func main()  {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
