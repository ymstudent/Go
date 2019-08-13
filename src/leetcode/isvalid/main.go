package main

import (
	"fmt"
)

func main()  {
	fmt.Println(isValid("()[]{}"))
}

/**
 * 思路：刚好最近看极客时间算法课，遇到类似的问题。遍历字符串，将左括号压入栈中，遇到右括号时从栈中取出一个左括号，看看是否匹配
 * ps:第一忘了单字符的情况，导致出错，真是处处有陷进啊
 * 耗时：0ms，内存：2.2mb，时间复杂度：O(n)
 */
func isValid(s string) bool {
	m_left := map[string]string{"(":")", "{":"}", "[":"]"}
	m_right := map[string]string{ ")":"(", "}":"{", "]":"["}
	var l []string
	for _, v := range s {
		if _, ok := m_left[string(v)]; ok {
			l = append(l, string(v)) //左括号压入栈中
		}
		if value, ok := m_right[string(v)]; ok {
			llen := len(l)
			if llen < 1 {
				return false
			}
			item := l[llen - 1]
			if item != value {
				return false
			}
			l = l[:llen - 1]
		}
	}
	//防止例如"["这样的单字符
	if len(l) > 0 {
		return false
	}
	return true
}
