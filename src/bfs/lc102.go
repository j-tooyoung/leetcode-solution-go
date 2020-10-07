package main

import "container/list"

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	l := list.New()
	l.PushBack(root)
	for l.Len() > 0 {
		size := l.Len()
		for i := 0; i < size; i++ {
			//l.Remove(i)

		}

	}
	return res
}

func main() {
	
}
