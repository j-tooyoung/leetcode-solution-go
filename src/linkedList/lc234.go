package main

// 回文链表
// 递归 todo
func isPalindrome(head *ListNode) bool {

	return false
}

// 链表反转 todo
func isPalindrome1(head *ListNode) bool {
	// fast,slow指针，将后半部分链表反转，逐一比较
	fast, slow := head, head
	for fast != nil && fast.Next != nil{
		fast = fast.Next.Next
		slow = slow.Next
	}
	return false
}

func isPalindrome2(head *ListNode) bool {
	var list []*ListNode
	for head!= nil {
		list =append(list, head)
		head = head.Next
	}
	i, j := 0, len(list) -1
	for i < j {
		if list[i].Val != list[j].Val {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	
}
