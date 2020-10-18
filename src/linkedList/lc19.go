package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummpy := &ListNode{Val: -1}
	dummpy.Next = head

	fast,slow := dummpy, dummpy
	cnt := 0
	for cnt < n &&fast!= nil {
		fast =fast.Next
		cnt++
	}
	if fast == nil {
		return nil
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	//
	slow.Next = slow.Next.Next
	return dummpy.Next
}