package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main()  {
	fmt.Println(comma("123342312312"))
	fmt.Println(comma2("-123123.123123"))
	fmt.Println(same("123", "321"))
}

//练习3.10 递归改循环，并用Buffer类型处理
func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	mod := n%3 //对3取余
	if mod > 0 { //大于0表示最左侧肯定不到3位，需要特殊处理
		buf.Write([]byte(s[:mod]+","))
	}
	for mod+3 < n { //循环加逗号
		buf.Write([]byte(s[mod:mod+3] + ","))
		mod += 3
	}
	if mod+3 == n { //最右侧3位特殊处理，因为最后面不需要加逗号
		buf.Write([]byte(s[mod:mod+3]))
	}

	return buf.String()
}

//练习3.11 支持浮点数以及带正负号的数字
func comma2(s string) string {
	//先剔除前面的符号
	start := ""
	if strings.HasPrefix(s, "-") || strings.HasPrefix(s, "+") {
		start = string(s[0])
		s = s[1:]
	}
	//然后剔除小数部分
	end := ""
	if strings.Contains(s, ".") {
		s_arr := strings.Split(s, ".")
		s , end = s_arr[0], "."+s_arr[1]
	}
	//最后处理整数部分
	s = comma(s)
	//拼接
	return strings.Join([]string{start, s, end}, "")
}

//练习3.12 判断2个字符串是否同文异构，即字符相同但顺序不同
func same(s1, s2 string) bool {
	//相等也不属于同文异构，不要忽略
	if s1 == s2 {
		return false
	}
	//长度不等
	if len(s1) != len(s2) {
		return false
	}
	//某个字符个数不等
	for _, v := range s1{
		//fmt.Printf("type:%T value:%v\n", v, v)
		//此处需要注意：v是int32，即rune类型，需要先转换为字符串
		if strings.Count(s1, string(v)) != strings.Count(s2, string(v)) {
			return false
		}
	}
	return true
}
