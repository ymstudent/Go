package main

import "fmt"

func main()  {
	node1 := ListNode{5, nil}
	node2 := ListNode{4, &node1}
	node3 := ListNode{4, &node2}
	node4 := ListNode{3, &node3}
	node5 := ListNode{3, &node4}
	node6 := ListNode{2, &node5}
	node7 := ListNode{1, &node6}
	fmt.Println(deleteDuplicates(&node7))
}

type ListNode struct {
	Val int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	res := &ListNode{Next:head} //为了方便第一个结点删除
	slow := res
	fast := head

	for fast != nil {
		if fast.Next != nil && fast.Val != fast.Next.Val {
			if slow.Next != nil {
				slow = fast
			} else {
				slow.Next = fast.Next
			}
		}
		fast = fast.Next
	}

	return res.Next
}