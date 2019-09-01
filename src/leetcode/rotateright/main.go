package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	test3 := ListNode{Val:3,Next:nil}
	test2 := ListNode{Val:2,Next:&test3}
	test1 := ListNode{Val:1,Next:&test2}
	res := rotaetRight(&test1, 2)
	fmt.Println(res.Val, res.Next.Val, res.Next.Next.Val)
}

func rotaetRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	oldList := head
	n := 1
	for oldList.Next != nil {
		oldList = oldList.Next
		n++
	}

	oldList.Next = head

	newList := head
	for i := 0; i < n - k % n - 1; i++ {
		newList = newList.Next
	}
	newHead := newList.Next
	newList.Next = nil
	return newHead
}
