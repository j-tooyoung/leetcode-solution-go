package main

// 222. 完全二叉树的节点个数
// https://leetcode-cn.com/problems/count-complete-tree-nodes/
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}
