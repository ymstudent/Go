package main

import "fmt"

func main()  {
	//关于指针
	a := 1
	b := a //这样其实是给b重新开辟了一块内存地址（而在PHP中，其实只是将b指向了a对应的内存地址，当重新给b赋值时才会开辟新的内存地址）
	fmt.Println(a, b)
	fmt.Println(&a, &b) //0xc00000a0b8 0xc00000a0d0

	//由下面的例子可以看到：*d实质上是c的别名，它和c指向了同一块内存地址
	c := 1
	d := &c
	e := *d //这里等价于 e:= c
	fmt.Println(&c, &d, &e, &*d) //0xc00000a110 0xc000006030 0xc00000a118 0xc00000a110

	f := newStruct()
	h := newStruct2()
	fmt.Println(f == h)


	f2()
	fmt.Println(&*global) //0xc00005a0d8, 变量x从f2中逃逸
}

var global *int

func f2()  {
	var x int
	x = 1
	fmt.Println(&x) //0xc00005a0d8
	global = &x
}

func newStruct() *struct{} {
	return new(struct{})
}

func newStruct2() *struct{}  {
	var a struct{}
	return &a
}
