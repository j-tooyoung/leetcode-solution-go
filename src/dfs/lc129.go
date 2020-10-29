package dfs

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 法1 dfs
func sumNumbers(root *TreeNode) int {
	var ans int
	var dfs func(root *TreeNode, tmp int) ;
	dfs = func(root *TreeNode, tmp int)  {
		if root == nil {
			return
		}
		tmp = tmp * 10 + root.Val
		if root.Left == nil && root.Right == nil {
			fmt.Println(tmp)
			ans =ans + tmp
			return
		}
		dfs(root.Left, tmp)
		dfs(root.Right, tmp)
	}
	dfs(root, 0)
	return ans
}

// 法2： 递归求左半部分和右半部分之和
func sumNumbers2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dfs2(root, 0)
}

func dfs2(root *TreeNode, tmp int) int {
	if root == nil {
		return 0
	}
	sum := tmp * 10 + root.Val
	if root.Left == nil && root.Right == nil {
		return sum
	}
	return dfs2(root.Left, sum) + dfs2(root.Right, sum)
}

// 法3： bfs
func sumNumbers3(root *TreeNode) int {
	return 0
}
