package main

import "fmt"

func main()  {
	//在Go中，双引号用来表示字符串sting，本质上是一个byte类型的数组
	//单引号用来表示rune类型，即int32类型
	s1 := "x"
	s2 := 'x'
	fmt.Printf("%t\n", s1 == s2) //错误：类型不同，不能直接比较
	//而反引号则用来创建原生的字符串字面量（可以看作带双引号的字节序列）,它可以由多行组成，不支持任何转义
	s3 := "x\n"
	s4 := `x\n`
	fmt.Printf("%t\n", s3 == s4) //此处永远为false，因为s3中的\n被转义了
}
