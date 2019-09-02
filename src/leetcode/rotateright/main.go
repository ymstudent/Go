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

//给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。

//思路：根据题目给出的2个示例，其实可以简化为3步
//1、将尾节点指针指向头节点，形成一个环形链表，并计算出链表长度n
//2、找到新的尾部，第n - k % n - 1个节点。新链表的头是第n - k % n个节点
//3、断开环，并返回新链表的头

/**
 * 耗时：4ms，内存：2.5mb，时间复杂度：O(n)
 */
func rotaetRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	//找到链表的尾部并计算链表长度
	oldTail := head
	n := 1
	for oldTail.Next != nil {
		oldTail = oldTail.Next
		n++
	}

	//将尾节点的指针指向头节点
	oldTail.Next = head

	//找到新的尾部，第n - k % n - 1个节点，断开环
	newTail := head
	for i := 0; i < n - k % n - 1; i++ {
		newTail = newTail.Next
	}
	newHead := newTail.Next
	newTail.Next = nil
	return newHead
}
