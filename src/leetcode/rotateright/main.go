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
	pointer := head
	for i := 0; i < k; i++ {
		for pointer.Next.Next != nil {
			pointer.Val = pointer.Next.Val
			pointer.Next = pointer.Next.Next
		}
		pointer.Next.Next = head
		pointer.Next = nil
		head = pointer
	}
	return pointer
}
