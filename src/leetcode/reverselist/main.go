package main

type ListNode struct {
	Val int
	Next *ListNode
}

func main()  {

}

//反转一个单链表。
//输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL

/**
 * 耗时：0ms，内存：2.6mb 时间复杂度：O(n)
 * 思路：在遍历链表时，将当前结点的指针指向上一个结点。但由于当前结点没有引用上一个结点，因此必须事先存储上一个结点。
 * 在更改指针指向之前，当需要用另一个变量将下一个结点存储下来。否则更改后就无法找到原链表中的下一个结点了
 */
func reverseList(head *ListNode) *ListNode {
	var prve *ListNode = nil
	for head != nil {
		next := head.Next	//用next保存下一个结点
		head.Next = prve    //将当前节点的指针指向前一个结点
		prve = head		    //当prev中保存的上一个结点替换为当前结点
		head = next			//移动到下一个结点
	}
	return prve
}
