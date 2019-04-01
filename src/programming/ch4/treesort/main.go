//利用二叉树实现插入排序
package main

import "fmt"

//定义一个二叉树
type tree struct {
	value int
	left, right *tree
}

func main()  {
	s := []int{7,1,3,2,8}
	Sort(s)
	fmt.Println(s)
}

//插入排序
func Sort(values []int) {
	var root *tree
	for _, v := range values{
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

//将元素按照顺序追加到values里面，然后返回结果slice
func appendValues(values []int, t *tree) []int {
	fmt.Println(values)
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree) //等价于返回 &tree{value: value}
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

