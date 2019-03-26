//因为slice包含了指向底层数组的指针，所以将一个slice传递给函数时，可以在函数内部修改底层数组的元素
// ---也就是说slice本质上是引用传递而不是值传递
package main

import "fmt"

func main()  {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a) //[5 4 3 2 1 0]原始值发生改变，引用传递
	reverseArr(a)
	fmt.Println(a) //[5 4 3 2 1 0]未能改变原始值，值传递

	//将一个slice左移n个元素。
	s := a[:]
	reverse(s[:2]) //第一步：反转前n个元素
	fmt.Println(s) //[4 5 3 2 1 0]
	reverse(s[2:]) //第二步：反转剩下的元素
	fmt.Println(s) //[4 5 0 1 2 3]
	reverse(s[:]) //第三步：反转整个slice
	fmt.Println(s) //[3 2 1 0 5 4]
}

func reverse(s []int)  {
	for i, j := 0, len(s) -1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverseArr(s [6]int)  {
	for i, j := 0, len(s) -1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	fmt.Println(s) //[0 1 2 3 4 5]
}
