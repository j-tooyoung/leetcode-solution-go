package main

import "sort"

// 148. 排序链表
// https://leetcode-cn.com/problems/sort-list/
// brute force
func sortList(head *ListNode) *ListNode {
	dummpy := new(ListNode)
	var list []int
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}
	sort.Ints(list)
	cur:= dummpy
	for _, val := range list {
		cur.Next = &ListNode{Val: val}
		cur =cur.Next
	}
	cur.Next = nil
	return dummpy.Next
}
func sortList1(head *ListNode) *ListNode {

	return nil
}
