package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println(lengthOfLastWord("a"))
}


//给定一个仅包含大小写字母和空格 ' ' 的字符串，返回其最后一个单词的长度。
//
//如果不存在最后一个单词，请返回 0 。
//
//说明：一个单词是指由字母组成，但不包含任何空格的字符串。


/**
 * 耗时：0ms，内存：2.2mb，复杂度：O(n)
 * 没什么技巧，只需注意一下去掉尾部的空格即可，从后往前遍历，遇到空格就返回
 */
func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	ret := 0
	lens := len(s)
	if lens == 0 {
		return 0
	}
	for i := lens - 1; i >= 0; i-- {
		if string(s[i]) == " " {
			return ret
		} else {
			ret ++
		}
	}
	return ret
}
