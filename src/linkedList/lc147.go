package main

// 147. 对链表进行插入排序
// https://leetcode-cn.com/problems/insertion-sort-list/
func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummpy := new(ListNode)
	dummpy.Next = head
	cur := head
	for cur.Next != nil {
		if cur.Next.Val < cur.Val {
			prev := dummpy
			val := cur.Next.Val
			node := cur.Next
			cur.Next = node.Next
			for prev.Next.Val < val {
				prev = prev.Next
			}

			tmp := prev.Next
			prev.Next = node
			node.Next = tmp
		}else{
			cur = cur.Next
		}
	}
	return dummpy.Next
}
