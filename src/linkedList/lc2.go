package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	v1, v2, carry := 0,0, 0

	node := &ListNode{-1, nil}
	root := node
	for  l1 != nil || l2 != nil {
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		} else {
			v1 = 0
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		} else {
			v2 = 0
		}
		carry = carry + v1 + v2
		node.Next = &ListNode{carry % 10, nil}
		node = node.Next
		carry /= 10
	}
	if carry > 0 {
		node.Next = &ListNode{carry % 10, nil}
	}
	return root.Next

}
func main() {

}
