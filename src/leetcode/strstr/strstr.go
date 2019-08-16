package main

import "fmt"

func main()  {
	fmt.Println(strStr("hello", "ll"))
}

//实现 strStr() 函数。
//给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

/**
 * 耗时：0ms, 内存：2.3mb 时间复杂度：O(n)
 * 思路：拿needle开头一个字符与haystack中的各个字符比较，如果有相等的，从haystack中截取对应长度的字符与needle比较 
 * 刚开始没有注意到needle比haystack长的情况，导致出现数组越界问题
 */
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	llen := len(needle)
	slen := len(haystack)
	for k, v := range haystack {
		if int(v) == int(needle[0]) {
			if k + llen > slen { //注意数组越界问题
				return -1
			}
			str := haystack[k:k+llen]
			if str == needle {
				return k
			}
		}
	}
	return -1
}