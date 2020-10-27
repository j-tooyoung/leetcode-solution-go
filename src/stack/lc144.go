package stack

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// stack模拟 morris遍历 todo
func preorderTraversal(root *TreeNode) []int {
	return nil
}

// 递归
func preorderTraversal1(root *TreeNode) []int {
	var res []int
	var dfs func(root *TreeNode)

	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return res
}
