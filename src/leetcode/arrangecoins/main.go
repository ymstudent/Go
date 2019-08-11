package main

import (
	"fmt"
	"math"
)

func main()  {
	fmt.Println(arrangeCoins2(5))
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


/**
 * 分析题目可以知道硬币的排列为差值为1的等差数列，利用等差数列求和公式可得：n = (x^2 + x)/2
 * 解这个二元一次方程式得：x = 根号下（2 * n + 0.25） - 0.5
 * 耗时：4ms，内存：2.3mb
 */
func arrangeCoins2(n int) int {
	x := 2 * float64(n) + 0.25
	ret := math.Sqrt(x) - 0.5
	return int(ret)
}


/**
 * 利用二分查找法求解
 */
func arrangeCoins3(n int) int {
	left, right := 1, n
	for left < right {
		mid := left + (right - left) / 2
		if cal(mid) == n {
			return mid
		} else if cal(mid) > n {
			if cal(mid - 1) < n {
				return mid - 1
			} else {
				right = mid - 1
			}
		} else {
			left = mid + 1
		}
	}

	return -1

}

func cal(num int) int {
	return (num * num + num) / 2
}