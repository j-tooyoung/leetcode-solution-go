package main

func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && slow != nil {
		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			return false
		}
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
