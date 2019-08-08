package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println(addStrings("1", "9"))
}

//给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。
//注意：
//num1 和num2 的长度都小于 5100.
//num1 和num2 都只包含数字 0-9.
//num1 和num2 都不包含任何前导零。
//你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。

/**
 * 耗时：12ms，内存消耗：6.7mb
 */
func addStrings(num1 string, num2 string) string {
	//思路与前面的二进制字符串相加一样，只不过这次是10进制
	//从后往前，逢10进1，字符相加减
	lenA := len(num1)
	lenB := len(num2)
	count := lenA
	if lenA > lenB {
		num2 = strings.Repeat("0", lenA - lenB) + num2
	} else {
		num1 = strings.Repeat("0", lenB - lenA) + num1
		count = lenB
	}

	var ret string
	flag := 0
	for i := count -1; i >= 0; i-- {
		t1 , t2 := 0, 0
		if num1[i] > 0 {
			t1 = int(num1[i] - '0')
		}
		if num2[i] > 0 {
			t2 = int(num2[i] - '0')
		}
		sum := t1 + t2 +flag
		if sum >= 10 {
			flag = 1
			ret = string(sum - 10 + '0') + ret
		} else {
			flag = 0
			ret = string(sum + '0') + ret
		}
	}
	if flag == 1 {
		ret = "1" + ret
	}
	return ret
}
