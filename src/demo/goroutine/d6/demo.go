package main

import (
	"demo/goroutine/d6/tree"
	"fmt"
	"sort"
)

//练习：等价二叉查找树
//1、实现Walk函数，2、测试Walk函数

//Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	rangeTree(t, ch)
	close(ch)
}

//遍历二叉树
func rangeTree(t *tree.Tree, ch chan int)  {
	if t == nil {
		return
	}
	rangeTree(t.Left, ch)
	ch <- t.Value
	rangeTree(t.Right, ch)
}

//3、用Walk实现Same函数来检测t1和t2是否存储了相同的值。4、测试Same函数
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	s1 := make([]int, 0)
	s2 := make([]int, 0)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for v := range ch1{
		s1 = append(s1, v)
	}
	for v := range ch2{
		s2 = append(s2, v)
	}

	if len(s1) != len(s2) {
		return false
	}

	sort.Ints(s1)
	sort.Ints(s2)
	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}
	return true
}

func main()  {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for v := range ch {
		fmt.Printf("get value: %v\n", v)
	}

	res := Same(tree.New(1), tree.New(1))
	fmt.Println(res)
}
