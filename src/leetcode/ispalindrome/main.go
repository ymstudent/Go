package main

import (
	"fmt"
	"strconv"
)

func main()  {
	fmt.Println(isPalindrome2(11211))
}

//判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
//进阶：你能不将整数转为字符串来解决这个问题吗？

/**
 * 耗时：24ms，内存：5mb
 */
func isPalindrome(x int) bool {
	//使用了前面的倒转整数的思路
	if x < 0 {
		return false
	}
	initial := x
	revers := 0
	for x != 0 {
		revers = revers * 10 + x % 10
		x = x / 10

	}
	if revers == initial {
		return true
	} else {
		return false
	}
}

/**
 * 将整数直接转化为字符串，然后开头与末尾比较，有一个不同，直接返回false
 * 耗时：20ms，内存：5.1mb
 */
func isPalindrome2(x int) bool {
	str := strconv.Itoa(x)
	length := len(str)
	for i, j := 0, length -1; i < j; i, j = i +1, j - 1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}
