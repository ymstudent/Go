package main

import "fmt"

func main()  {
	//切片则为数组元素提供动态大小的、灵活的视角
	//类型 []T 表示一个元素类型为 T 的切片。
	//切片通过两个下标来界定，即一个上界和一个下界，二者以冒号分隔，它会选择一个半开区间，包括第一个元素，但排除最后一个元素。
	a := [5]int{1,2,3,4,5}
	s := a[1:3]
	fmt.Println(s)
	fmt.Printf("type: %T\n", s)
	//切片并不存储任何数据，它只是描述了底层数组中的一段。
	//更改切片的元素会修改其底层数组中对应的元素。
	//与它共享底层数组的切片都会观测到这些修改。
	s[0] = 10
	fmt.Println(a)
	fmt.Println(s)
	//下面这种做法会先创建一个[10]int数组，然后构建一个引用了它的切片
	s2 := []int{1,2,3,4,5,6,7,8,9,10}
	//在进行切片时，你可以利用它的默认行为来忽略上下界。这2个切片是相等的
	s2 = s2[0:4]
	s2 = s2[:4]
	//切片拥有 长度 和 容量。切片的长度就是它所包含的元素个数。
	//切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。
	//切片 s 的长度和容量可通过表达式 len(s) 和 cap(s) 来获取。
	s3 := []int{2, 3, 5, 7, 11, 13}
	printSlice(s3)
	s3 = s3[:0] // 截取切片使其长度为 0
	printSlice(s3)
	s3 = s3[:4] // 拓展其长度
	printSlice(s)
	s3 = s3[2:] // 舍弃前两个值
	printSlice(s3)
	//切片的零值是 nil。nil 切片的长度和容量为 0 且没有底层数组。
	var s4 []int
	printSlice(s4)
	if s4 == nil {
		fmt.Println("nil!")
	}
	//使用make方法创建切片，函数会分配一个元素为零值的数组并返回一个引用了它的切片。第一个参数指定类型，第二个参数指定长度，第三个参数指定容量
	s5 := make([]int, 5, 5) //容量参数可选，不能小于长度
	printSlice(s5)
	//使用append向切片中追加元素
	var s6 []int
	printSlice(s6)
	//添加一个空切片
	s6 = append(s6, 0)
	printSlice(s6)
	s6 = append(s6,1)
	printSlice(s6)
	s6 = append(s6,2,3,4,5)
	printSlice(s6)
}

func printSlice(s []int)  {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
