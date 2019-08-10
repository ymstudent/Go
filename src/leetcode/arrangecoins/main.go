package main

import (
	"fmt"
)

func main()  {
	fmt.Println(arrangeCoins(10))
}

//你总共有 n 枚硬币，你需要将它们摆成一个阶梯形状，第 k 行就必须正好有 k 枚硬币。
//给定一个数字 n，找出可形成完整阶梯行的总行数。

/**
 * 耗时：12ms，内存：2.2mb
 */
func arrangeCoins(n int) int {
	//暴力解法：逐行减，到那一行不够，直接返回行数-1
	i := 1
	for {
		n = n - i
		if n < 0 {
			return i - 1
		}
		i++
	}
}
