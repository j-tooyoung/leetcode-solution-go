package main

import "math"

var ans int
//var pre *TreeNode
var pre int

func Min(a,b int) int{
	if a < b {
		return a
	}else {
		return b
	}
}

func Abs(a, b int) int  {
	if a < b {
		return b - a
	}else {
		return a - b
	}
}
func abs(x int) int  {
	if x < 0 {
		return -x
	}else {
		return x
	}
}

// 内部函数写法
func getMinimumDifference1(root *TreeNode) int {
	res, prev := math.MaxInt64, -1
	var dfs3 func( *TreeNode)
	dfs3 = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs3(node.Left)
		// 减小函数调用，速度大幅提升
		if prev != -1 && node.Val -prev < res{
			res =node.Val-prev
		}
		prev= node.Val
		dfs3(node.Right)
	}
	dfs3(root)
	return res

}

func getMinimumDifference(root *TreeNode) int {
	//ans := int(^uint(0) >> 1)
	ans := math.MaxInt32
	pre = -1
	//var list []int
	//dfs2(root,&list)
	//for i := 0; i + 1 < len(list); i++ {
	//	ans = Min(ans, Abs(list[i+1],list[i]))
	//}
	dfs1(root)
	return ans
}

// 有问题 fixme
func dfs1(root *TreeNode) {
	if root == nil {
		return
	}
	dfs1(root.Left)
	if pre!= -1 {
		ans = Min(ans, abs(root.Val-pre))
	}
	pre = root.Val
	dfs1(root.Right)
}

func dfs2(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}
	dfs2(root.Left, list)
	*list = append(*list, root.Val)
	dfs2(root.Right, list)
}




