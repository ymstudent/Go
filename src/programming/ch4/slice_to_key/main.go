//将一个字符串slice转换为一个适合做map键的字符串。
//该方法适用于任何不可直接比较的键类型
package main

import "fmt"

var m = make(map[string]int)

func main()  {
	s := []string{"a"}
	for i:=0; i< 10; i++ {
		Add(s)
	}
	res := Count(s)
	fmt.Println(res)
}

//第一步：定义一个帮助函数k将每一个键都映射到字符串，当且仅当x和y相等时，我们才认为k(x)==k(y)
//然后创建一个map，map的键是字符串类型，在每个键元素被访问的时候，调用这个帮助函数
func k(list []string) string {
	return fmt.Sprintf("%q", list) //将一个字符串slice转换为一个字符串
}

func Add(list []string) {
	m[k(list)]++
}

func Count(list []string) int {
	return m[k(list)]
}
