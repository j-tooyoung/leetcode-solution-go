package main

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast!=nil && slow != nil {
		if fast.Next!= nil {
			fast = fast.Next.Next
		}else {
			return nil
		}
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	if fast == nil {
		return nil
	}
	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

func main() {
	
}
