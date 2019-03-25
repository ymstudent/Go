package main

import "fmt"

func main()  {
	m := make(map[string]int)
	//在映射m中插入或者修改元素 m[key] = elem
	m["Bell"] = 1 //插入
	fmt.Println(m)
	m["Bell"] = 2 //修改
	fmt.Println(m)

	//获取元素：elem := m[key]
	v := m["Bell"]
	fmt.Println(v)

	//删除元素 delete(m, key)
	delete(m, "Bell")
	fmt.Println(m)

	//通过双赋值检测某个键是否存在。elem, ok := m[key],存在则ok=true，elem为其对应值；不存在则ok=false，elem为其映射元素类型的零值
	elem, ok := m["Bell"]
	fmt.Println(elem, ok)

}
