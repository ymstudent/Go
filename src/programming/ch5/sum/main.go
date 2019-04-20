//变长函数
package main

import (
	"fmt"
	"os"
)

func main()  {
	fmt.Println(sum())
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1,2,3,4))
	s := []int{1,2,3,4}
	fmt.Println(sum(s...)) //当实参为slice时如何调用变长函数：在最后一个参数前放一个省略号

	linenum , name := 12, "count"
	errorf(linenum, "undefined %s", name)
}

func sum(vals ...int) int { //在类型名称声明之前使用省略号，表示声明一个变长函数，其可以接受任意数目的参数
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func errorf(linenum int, format string, args ...interface{}) { //变长函数通常用于格式化字符串， interface{}类型意味着函数的最后一个参数可以接受任何值
	fmt.Fprintf(os.Stderr, "line %d: ", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}
