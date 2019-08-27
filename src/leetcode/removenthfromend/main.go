package main

func main()  {

}

//给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

type ListNode struct {
	Val int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}
	//添加一个哨兵节点。（ps：为什么要添加一个哨兵节点:简化某些极端情况。如删除头节点和只包含一个节点的链表）
	//eg:链表：1—>2，n=2，在不添加哨兵节点的情况下，如何删除头节点？
	//普通节点的删除操作：node.Next = node.Next.Next不适用于头节点删除。
	node := &ListNode{Next:head}
	pointer, pointer2, length := node, node, 1
	//统计链表长度
	for pointer.Next != nil {
		pointer = pointer.Next
		length++
	}
	index := length - n
	for i := 1; i <= length; i++ {
		if i == index {
			if pointer2.Next == nil {break} //
			pointer2.Next = pointer2.Next.Next
			break
		} else {
			pointer2 = pointer2.Next
		}
	}
	return node.Next
}
