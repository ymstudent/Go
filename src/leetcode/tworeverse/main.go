package main

import (
	"fmt"
	"math"
)

func main()  {
	fmt.Println(126 /  10)
	fmt.Println(reverse(1234))
}

//给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

/**
 * 耗时：4ms，内存：2.2mb
 */
func reverse(x int) int {
	//go除法运算（/）的行为取决于操作数是否都为整型，整数相除，商会舍弃小数部分。例如：5.0/4.0=1.25, 5/4=1
	//取模余数（%）的正负号总是与被除数一致。例如-5%3与-5%-3都等于-2
	ret := 0
	for x != 0 {
		pop := x % 10
		x = x / 10
		ret = ret * 10 + pop
		if ret < math.MinInt32 || ret > math.MaxInt32 {
			return 0
		}
	}
	return ret
}
