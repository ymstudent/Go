package main

import "fmt"

func main()  {
	//注意字符串s的字串操作和对字节slice做slice操作的区别
	//区别就是字符串s的字串操作返回的是一个字符串，而对字节slice做slice操作返回的是一个slice
	s := "hello" //字符串的底层其实是一个byte数组
	byte_slice := []byte(s) //[104 101 108 108 111]
	fmt.Println(byte_slice)
	sub_s := s[1:3] //求字符串s的字串操作
	sub_slice := byte_slice[1:3] //对字节slice做slice操作
	fmt.Println(sub_s, sub_slice)
}
