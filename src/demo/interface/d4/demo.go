package main

import "fmt"

//底层值为nil的接口值
//即使接口内的具体值为nil，方法仍然会被nil接收者调用
//在一些语言中，这会触发一个空指针异常，但在Go中通常会写一些方法来优雅的处理它（如本例中的M方法）
//注意：保存了nil具体值得接口其自身并不为nil

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M()  {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main()  {
	var i I

	var t *T
	i = t //此时结构体T未赋值，所以其指针指向的底层值（*T）为nil（默认零值）
	describe(i)
	i.M()

	i = &T{"Hello World!"}
	describe(i)
	i.M()
}

func describe(i I)  {
	fmt.Printf("(%v, %T)\n", i, i)
}

