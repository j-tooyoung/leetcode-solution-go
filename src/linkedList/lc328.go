package main

/**
奇偶链表
*/
// 暴力解法
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var list []*ListNode
	for head != nil {
		list = append(list, head)
		head = head.Next
	}
	idx := 0
	for ; idx+2 < len(list); idx += 2 {
		list[idx].Next = list[idx+2]
	}
	j := 1
	for ; j+2 < len(list); j += 2 {
		list[j].Next = list[j+2]
	}
	list[j].Next = nil
	list[idx].Next = list[1]
	return list[0]
}

// 原地算法
func oddEvenList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	odd, even := head, head.Next
	second := even
	for odd.Next != nil && even.Next != nil {
		odd.Next = even.Next
		if odd.Next != nil {
			even.Next = odd.Next.Next
		}else {
			break
		}
		odd = odd.Next
		even = even.Next
	}
	odd.Next = second
	return head
}
