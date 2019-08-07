package main

import "fmt"

func main()  {
	fmt.Println(addDigits(22))
}

//给定一个非负整数 num，反复将各个位上的数字相加，直到结果为一位数。

/**
 * 耗时：0ms，内存：2.2mb
 */
func addDigits(num int) int {
	if num < 10 {
		return num
	}
	//各位相加，大于10继续循环
	for {
		ret := 0
		for num != 0 {
			tmp := num % 10
			num /= 10
			ret += tmp
		}
		if ret < 10 {
			return ret
		}
		num = ret
	}
}

//进阶：你可以不使用循环或者递归，且在 O(1) 时间复杂度内解决这个问题吗？
//思路：假设一个三位整数：n = 100*a + 10*b + c，按上面的暴力解法的思路循环一次可以得到 addn = a+b+c
//两者的差值：n - addn = 99*a + 9*b = 9 * (11*a + b)，差值可以被9整除，说明每次循环缩小9的倍数
//所以我们直接将num对9取余，如果不为0，则返回余数，如果为0，则返回9
//这个思路是看leetcode上的解答写的，实在太牛逼了
func addDigits2(num int) int {
	if num > 9 {
		num = num % 9
		if num == 0 {
			return 9
		}
	}
	return num
}
