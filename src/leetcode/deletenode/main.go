package main

type ListNode struct {
	Val int
	Next *ListNode
}

func main()  {

}


//请编写一个函数，使其可以删除某个链表中给定的（非末尾）节点，你将只被给定要求被删除的节点。
//现有一个链表 -- head = [4,5,1,9]，它可以表示为4->5->1->9 
//说明:
//链表至少包含两个节点。 
//链表中所有节点的值都是唯一的。 
//给定的节点为非末尾节点并且一定是链表中的一个有效节点。 
//不要从你的函数中返回任何结果。


/** 
 * 这题比较坑，我还想怎么只有一个参数，起码应该再给一个参数，告诉我删除的是啥啊。 
 * 还来才想清楚。他是已经指定了链表；head = [4,5,1,9]，那个node，是要删除的元素。所以这题的意思是删除链表的当前节点 
 * 阅读理解啊 
 */
func deleteNode(node *ListNode)  {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
