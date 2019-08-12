package main

import "fmt"

func main()  {
	fmt.Println(romanToInt("IX"))
}

//罗马数字转整数

/**
 * 耗时：8ms，内存：3mb
 */
func romanToInt(s string) int {
	m := map[string]int{
		"I" : 1,
		"V" : 5,
		"X" : 10,
		"L" : 50,
		"C" : 100,
		"D" : 500,
		"M" : 1000,
	}

	llen := len(s)
	ret := 0
	for i := 0; i < llen; i ++ {
		n1 := m[string(s[i])]
		if i + 1 >= llen {
			ret += n1
			break
		}
		n2 := m[string(s[i+1])]
		if n1 < n2 {
			ret += n2 - n1
			i++
		} else {
			ret += n1
		}
	}
	return ret
}
