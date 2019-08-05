package main

import "fmt"

func main()  {
	fmt.Println(isOneBitCharacter([]int{1,0,0,1,1,0,0,1,0}))
}

//有两种特殊字符。第一种字符可以用一比特0来表示。第二种字符可以用两比特(10 或 11)来表示。
//现给一个由若干比特组成的字符串。问最后一个字符是否必定为一个一比特字符。给定的字符串总是由0结束。

/**
 * 耗时：4ms， 内存：2.7mb
 */
func isOneBitCharacter(bits []int) bool {
	//1、如果数组长度为1，必定是1bit字符，直接返回true
	//2、如果某一位为1，那么必定可以和它的下一位组成一个2bit字符，所以直接i++，移到下下位判断
	length := len(bits)
	if length == 1 {
		return true
	}
	var isone bool
	for i := 0; i < length; i++ {
		if bits[i] == 1 {
			i++
			isone = false
		} else {
			isone = true
		}
	}
	return isone
}