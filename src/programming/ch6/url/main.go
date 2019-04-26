package main

import "fmt"

//声明一个字符串到字符串列表的映射
type Values map[string][]string

//Get返回第一个具有给定key的值，不存在则返回空字符串
func (v Values) Get(key string) string {
	if vs := v[key]; len(vs) > 0 {
		return vs[0]
	}
	return ""
}

//Add添加一个键值到给定key的列表中
func (v Values) Add(key, value string)  {
	v[key] = append(v[key], value)
}

// 方法和函数一样，对引用本身作任何改变，都不会在调用者身上产生作用，所以下面的m即使调用Setnil，m的值也未改变
func (v Values) Setnil() {
	v = nil
}

func (v Values) Setnil2(key string) {
	v[key] = nil
}

func main()  {
	m := Values{"lang":{"en"}}
	m.Add("item", "1")
	m.Add("item", "2")
	fmt.Println(m.Get("lang")) //"en"
	fmt.Println(m.Get("q")) //""
	fmt.Println(m.Get("item")) //"1" (第一个值)
	fmt.Println(m["item"]) //"[1, 2]" (直接访问map)
	m.Setnil()
	fmt.Println(m) //[item:[1 2] lang:[en]]
	m.Setnil2("lang")
	fmt.Println(m) //[lang:[] item:[1 2]]

	m = nil
	fmt.Println(m.Get("item")) //""
	fmt.Println(m) //[]
	//m.Add("item", "3") //宕机：赋值给nil类型
}

