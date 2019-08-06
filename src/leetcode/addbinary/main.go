package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println(addBinary("11", "1"))
}

//给定两个二进制字符串，返回他们的和（用二进制表示）。
//输入为非空字符串且只包含数字 1 和 0。


/**
 * 耗时：0ms, 内存：2.6mb
 */
func addBinary(a string, b string) string {
	//为较短的字符串进行补0
	lenA := len(a)
	lenB := len(b)
	count := lenA
	if lenA > lenB {
		b = strings.Repeat("0", lenA - lenB) + b
	} else {
		a = strings.Repeat("0", lenB - lenA) + a
		count = lenB
	}
	//从最后一位算起
	var ret string
	flag := 0
	for i := count - 1; i >= 0; i-- {
		t1 , t2 := 0, 0
		if a[i] > '0' {
			//字符相减。利用ASCII码，字符在内部都用数字表示，如'1' = 49， '0' = 48，字符的加减，大小的比较实际上都是内部数字的加减，大小比较
			//所以：'1' - '0' = 1代表的ASCII码，我们是有int()即可将其转化位整数1
			t1 = int(a[i] - '0')
		}
		if b[i] > '0' {
			t2 = int(b[i] - '0')
		}
		//fmt.Println(a[i], b[i], t1, t2)
		sum := t1 + t2 + flag
		switch sum {
			case 3: flag = 1; ret = "1" + ret
			case 2: flag = 1; ret = "0" + ret
			case 1: flag = 0; ret = "1" + ret
			case 0: flag = 0; ret = "0" + ret
		}
	}
	//判断最后是否需要进位
	if flag == 1{
		ret = "1" + ret
	}
	return ret
}