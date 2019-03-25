/**
 * 知识点：通过bytes包提供的Buffer类型高效处理字节slice
 */
package main

import (
	"bytes"
	"fmt"
)

func main()  {
	fmt.Println(intsToString([]int{1,2,3})) //"[1,2,3]"
}

//intsToString与fmt.Sprint(values) 类似，但插入了逗号
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf,"%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
