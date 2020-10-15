package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// fixme ?

func connect(root *Node) *Node {
	if root == nil || root.Left == nil {
		return root
	}
	root.Left.Next = root.Right
	if root.Next != nil {
		root.Right.Next = root.Next.Left
	}
	connect(root.Left)
	connect(root.Right)
	return root
}

// 法2 遍历
// 法3 bfs
func connect1(root *Node) *Node {
	if root == nil {
		return root
	}
	// pre, cur 下一层，上一层
	//dummpy := &Node{Val: -1}
	//pre, cur := dummpy, root
	var pre, cur *Node
	cur = root
	for cur != nil {
		dummpy := &Node{Val: -1}
		pre = dummpy
		pre.Next = cur.Left
		pre = pre.Next
		pre.Next = cur.Right
		pre = pre.Next

		cur = cur.Next
		if cur == nil {
			cur = dummpy.Next
		}
	}
	return root
}
