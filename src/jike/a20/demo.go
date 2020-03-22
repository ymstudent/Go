package main

import "fmt"

//从panic引发到程序终止运行的大致过程
//某行代码引发panic，这时初始的panic详情会被建立起来，
//并且程序的控制权会从改行代码转移到其所属的函数的那行代码上，也就是调用栈的上一级，
//这就代表这行代码所属函数终止执行。随后控制器沿着栈调用的反方向一级一级的传播至顶端，
//也就是最外层的main函数，最终控制器这go运行时系统回收，此时程序崩溃并终止运行
//在控制权传播的过程中，panic详情会被逐渐积累并完善，并在程序终止之前打印出来

func main()  {
	fmt.Println("Enter function main.")
	caller1()
	fmt.Println("Exit function main.")
}

func caller1()  {
	fmt.Println("Enter function caller1.")
	caller2()
	fmt.Println("Exit function caller1.")
}

func caller2()  {
	//如何施加对panic的保护措施，防止程序崩溃
	defer func() {
		fmt.Println("Enter defer.")
		p := recover(); if p != nil {
			fmt.Printf("panic: %s\n", p)
		}
		fmt.Println("Exit defer.")
	}()

	fmt.Println("Enter function caller2.")
	s1 := []int{1,2,3,4,5}
	e5 := s1[5]
	_ = e5
	fmt.Println("Exit function caller2.")
}


