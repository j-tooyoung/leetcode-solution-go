package main

/**
24. 两两交换链表中的节点

*/
// 遍历
func swapPairs(head *ListNode) *ListNode {
	dummpy := &ListNode{Val: -1}
	dummpy.Next = head
	pre, cur := dummpy, head
	for cur != nil {
		var tmp *ListNode
		if cur.Next != nil {
			tmp = cur.Next
		}else{
			break
		}
		pre.Next = tmp
		cur.Next = tmp.Next
		tmp.Next = cur

		// pre,cur继续
		//pre = cur
		//cur = cur.Next
		pre,cur = cur,cur.Next

	}
	return dummpy.Next
}

func swapPairs1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairs1(newHead.Next)
	newHead.Next = head
	return newHead
}

func main() {

}
