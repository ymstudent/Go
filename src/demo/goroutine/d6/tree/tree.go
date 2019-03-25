package tree

import (
	"fmt"
	"math/rand"
)

// A Tree is a binary tree with integer values.
type Tree struct {
	Left *Tree
	//注意：这里的V需要大写，因为在在 Go 中，如果一个名字以大写字母开头，那么它就是已导出的。
	// 导入一个包时，你只能引用其中已导出的名字。任何“未导出”的名字在该包外均无法访问。
	//开始我忘记了这点，导致怎么也没有发现value
	Value int
	Right *Tree
}

//// New returns a new, random binary tree holding the values k, 2k, ..., 10k.
func New(k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(10){
		t = insert(t, (1+v)*k)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}

	if v < t.Value {
		t.Left = insert(t.Left, v)
	}else {
		t.Right = insert(t.Right, v)
	}

	return t
}

func (t *Tree) String() string {
	if t == nil {
		return "()"
	}

	s := ""
	if t.Left != nil {
		s += t.Left.String() + " "
	}
	s += fmt.Sprint(t.Value)
	if t.Right != nil {
		s += " " + t.Right.String()
	}

	return "(" + s + ")"
}
