package main

func inorderTraversal(root *TreeNode) []int {
	var res []int
	dfs(root,&res)
	return res
}

func dfs(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	dfs(root.Left, res)
	*res = append(*res, root.Val)
	dfs(root.Right, res)
}

func main() {
	
}
