package main

import (
	"fmt"
)

func main()  {
	fmt.Println(addToArrayForm2([]int{0}, 23))
}

//对于非负整数 X 而言，X 的数组形式是每位数字按从左到右的顺序形成的数组。例如，如果 X = 1231，那么其数组形式为 [1,2,3,1]。
//给定非负整数 X 的数组形式 A，返回整数 X+K 的数组形式。

/**
 * 耗时：500ms，内存：8.9mb
 */
func addToArrayForm(A []int, K int) []int {
	//分解K
	var s []int
	for K != 0 {
		s = append(s, K % 10)
		K /= 10
	}
	//倒转s
	for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
		s[i], s[j] = s[j], s[i]
	}
	//较短的补0
	lenA := len(A)
	lenS := len(s)
	if lenA > lenS {
		count := lenA - lenS
		var newS []int
		for i := count; i > 0; i-- {
			newS = append(newS, 0)
		}
		for _,v := range s {
			newS = append(newS, v)
		}
		s = newS
	} else {
		count := lenS - lenA
		var newA []int
		for i := count; i > 0; i-- {
			newA = append(newA, 0)
		}
		for _,v := range A {
			newA = append(newA, v)
		}
		A = newA
		lenA = len(A)
	}
	//逐位相加
	tmp := 0
	flag := 0
	var ret []int
	for i := lenA - 1; i >= 0; i-- {
		tmp = s[i] + A[i] + flag
		if tmp > 9 {
			flag = 1
			tmp = tmp - 10
		} else {
			flag = 0
		}
		ret = append(ret, tmp)
	}

	if flag == 1 {
		ret = append(ret, 1)
	}
	//倒转结果
	for i, j := 0, len(ret) - 1; i < j; i, j = i + 1, j - 1 {
		ret[i], ret[j] = ret[j], ret[i]
	}
	return ret
}

/**
 * 耗时：504ms, 内存9.1mb
 */
func addToArrayForm2(A []int, K int) []int {
	length := len(A)
	tmp := K
	var s []int
	for tmp > 0 || length > 0{
		if length > 0 {
			tmp = tmp + A[length - 1]
		}
		s = append(s, tmp % 10)
		length--
		tmp /= 10
	}
	for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
