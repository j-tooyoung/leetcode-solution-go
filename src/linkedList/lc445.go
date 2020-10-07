package main

/**
445. 两数相加 II
https://leetcode-cn.com/problems/add-two-numbers-ii/
*/

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	var list1 []int
	var list2 []int
	for l1 != nil {
		//list1 = append(list1, l1.Val)
		//l1 = l1.Next
		list1,l1 = append(list1,l1.Val),l1.Next
	}
	for l2 != nil {
		list2 = append(list2, l2.Val)
		l2 = l2.Next
	}
	i, j, carry := len(list1)-1, len(list2)-1, 0
	var node *ListNode
	for i >= 0 || j >= 0 || carry > 0 {
		a, b, sum := 0, 0, 0
		if i >= 0 {
			a = list1[i]
		}
		if j >= 0 {
			b = list2[j]
		}
		sum = a + b + carry
		carry = sum / 10
		//使用头插法插入节点
		tmp := &ListNode{Val: sum % 10}
		tmp.Next = node
		node = tmp
		i--
		j--
	}

	return node
}

func main() {

}
