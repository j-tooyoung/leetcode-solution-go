package main

// https://leetcode-cn.com/problems/reorder-list/submissions/
// 递归
func reorderList(head *ListNode)  {

}

func reorderList1(head *ListNode)  {
	if head == nil || head.Next == nil {
		return
	}
	var list []*ListNode
	cur := head
	for cur!= nil {
		list = append(list, cur)
		cur = cur.Next
	}
	//cur = head
	i,j := 0, len(list) -1
	for i < j {
		list[i].Next = list[j]
		list[j].Next = list[i + 1]
		i++
		j--
	}
	list[i].Next = nil

}
