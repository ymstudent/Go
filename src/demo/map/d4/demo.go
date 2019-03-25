package main

import (
	"fmt"
	"strings"
)

//练习:映射 实现 WordCount。它应当返回一个映射，其中包含字符串 s 中每个“单词”的个数。
func WordCount(s string) map[string]int {
	slice := strings.Fields(s)
	m := make(map[string]int)
	for _,v := range slice {
		elem, ok := m[v]
		if ok {
			m[v] = elem + 1
		}else {
			m[v] = 1
		}
	}
	return m
}

func main()  {
	m := WordCount("a a b s b b")
	fmt.Println(m)
}
