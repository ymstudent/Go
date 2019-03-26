//append函数的简单模拟
package main

import "fmt"

func main()  {
	var x, y []int
	for i := 0; i < 10 ; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}

func appendInt(x []int, y int) []int {
	var z []int
	fmt.Println(len(x), cap(x))
	zlen := len(x) + 1
	if zlen <= cap(x) {
		//slice仍有增长空间，扩展slice内容
		fmt.Println(zlen)
		z = x[:zlen]
	} else {
		//slice已无空间，为它重新分配一个底层数组
		//为了达到分摊线性复杂性，容量扩容一倍
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2*len(x)
		}
		fmt.Println(zlen, zcap)
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}