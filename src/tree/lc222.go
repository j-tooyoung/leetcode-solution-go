package main

// 222. 完全二叉树的节点个数
// https://leetcode-cn.com/problems/count-complete-tree-nodes/
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}

// 二分+位运算
func countNodes1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := depth(root.Left)
	rightHeight := depth(root.Right)
	if leftHeight == rightHeight {
		return countNodes1(root.Right) + (1 << leftHeight)
	}else{
		return countNodes1(root.Left) + (1<< rightHeight)
	}
}

func depth(node *TreeNode) int {
	level := 0
	for node != nil {
		node = node.Left
		level++
	}
	return level
	//if node == nil {
	//	return 0
	//}
	//left := depth(node.Left)
	//right := depth(node.Right)
	//if left <right {
	//	return right + 1
	//}else{
	//	return left + 1
	//}
}
