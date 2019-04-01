package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main()  {
	a := [...]int{0, 1, 2, 3, 4, 5}
	//reverse(&a)
	//fmt.Println(a)

	s := a[:]
	res := rotate(s, 2)
	fmt.Println(res)

	//s2 := []string{"a", "a", "b", "c", "c", "c", "d"}
	//s2 = noRepeat(s2)
	//fmt.Println(s2)

	//[]rune和[]byte在书中都叫字节slice，在处理英文字符串时看不处理区别，但我们转换一个中文字符就明白了
	s3 := "a    b   c"
	fmt.Println([]rune(s3), []byte(s3)) //[97 32 32 32 32 98 32 32 32 99] [97 32 32 32 32 98 32 32 32 99]
	s4 := "我   爱   你"
	fmt.Println([]rune(s4), []byte(s4)) //[25105 32 32 32 29233 32 32 32 20320] [230 136 145 32 32 32 231 136 177 32 32 32 228 189 160]
	res2 := noContinueEmpty([]rune(s4))
	fmt.Println(res2)

	wordfreq()
}

//4.3重写reverse函数，使用数组指针作为参数
func reverse(p *[6]int)  {
	fmt.Printf("type=%T\tvalue=%v\n", p, p)
	fmt.Printf("type=%T\tvalue=%v\n", *p, *p)
	for i, j := 0, len(p) - 1; i < j; i, j = i+1, j-1 {
		(*p)[i], (*p)[j] = (*p)[j], (*p)[i] //此处可以写成：p[i], p[j] = p[j], p[i]
											//因为为了方便，此处编译器会将p[i]解释为(*p)[i]
	}
}

//4.4编写一个rotate函数，实现只遍历一次就可以完成元素旋转
func rotate(s []int, n int) []int {
	s2 := s[n:]
	for i := n -1 ; i >= 0; i-- {
		s2 = append(s2, s[i])
	}
	return s2
}

//4.5编写一个就地处理函数，用于除去[]string slice中相邻的重复字符串
func noRepeat(s []string) []string {
	if len(s) < 1 {
		return s
	}
	n := 0
	current := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] != current {
			n++
			s[n] = s[i]
			current = s[i]
			fmt.Println(i, n, s)
		}
	}
	return s[:n+1]
}

//4.6编写一个就地处理函数，将一个UTF-8编码的字节slice中所有相邻的Unicode空白字符缩减为一个ASCⅡ空白字符（即将多个连续空格换成单个空格）
func noContinueEmpty(s []rune) []rune {
	n := 0
	current := s[0]
	for i := 1; i < len(s); i++ {
		//当连续位相等时，只有加一个不为空的判断就行
		if s[i] != current || !unicode.IsSpace(current) {
			n++
			s[n] = s[i]
			current = s[i]
			fmt.Println(i, n, s)
		}
	}
	return s[:n+1]
}

//4.7修改reverse函数，在不重新分配内存的情况下。来反转一个UTF-8编码的字符串中的字符元素，传入的参数是该字符串对应的字节slice类型
//TODO
//没想出来，不能重新分配内存就表示slice指向的底层数组不能变

//4.8修改charcount的代码来统计字母、数字和其他在Unicode分类中的字符数量
func charcount()  {
	in := bufio.NewReader(os.Stdin)
	var letter, number int
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if unicode.IsLetter(r) { //字母
			letter++
		}
		if unicode.IsNumber(r) { //数字
			number++
		}
		//其余的类型可以参考graphic.go包，这里就不一一写出
	}
	fmt.Printf("letter: %d\n", letter)
	fmt.Printf("number: %d\n", number)
}

//4.9编写一个wordfreq来汇总输入文本文件中的每个单词出现的次数
func wordfreq()  {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanWords) //将文本行按单词分割
	counts := make(map[string]int)
	for scanner.Scan() {
		counts[scanner.Text()]++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}
	for i, v := range counts {
		fmt.Printf("%s\t%d\n", i, v)
	}
}
