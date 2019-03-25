/**
 * 输入一个表示整数的字符串，从右侧开始，每隔3个数字后面就插入一个“,”，形如“12,345”
 */
package main

import "fmt"

func main()  {
	/*字符串可以与字节slice相互转换
	a := "abc"
	b := []byte(a)
	a2 := string(b)
	fmt.Println(b, a2)*/
	s := "21323123424312123"
	s = comma(s)
	fmt.Println(s)
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3])+","+s[n-3:]
}
