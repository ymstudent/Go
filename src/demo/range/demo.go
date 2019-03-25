package main

import "fmt"

//练习：切片
func Pic(dx, dy int) [][]uint8 {
	s := make([][]uint8, dy)

	for i := range s {
		s[i] = make([]uint8, dx)
	}

	return s
}


func main()  {
	s := make([]int, 0)
	fmt.Println(s)
	s = append(s,1,2,3,4)
	fmt.Println(s)
	//for 循环的 range 形式可遍历切片或映射。
	for i,v := range s {
		fmt.Printf("key: %d => value: %d\n", i, v)
	}
	//可以将下标或值赋予 _ 来忽略它。
	for k := range s{
		fmt.Printf("key => %d\n", k)
	}

	for _, v := range s{
		fmt.Printf("value => %d\n", v)
	}

	power := make([]int,10)

	for i := range power {
		power[i] = 1 << uint(i)
	}

	for _, value := range power {
		fmt.Printf("%d\n", value)
	}

}
