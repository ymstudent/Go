//统计输入中Unicode代码点出现次数的程序：使用map来追踪每个字符出现的次数
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main()  {
	counts := make(map[rune]int) //Unicode字符数量
	var utflen [utf8.UTFMax + 1]int //UTF-8编码长度
	invalid := 0 //非法UTF-8字符数量

	in := bufio.NewReader(os.Stdin)

	for {
		r, n, err := in.ReadRune() //返回rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
