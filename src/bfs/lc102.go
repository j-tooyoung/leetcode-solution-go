package main

import "container/list"

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root ==nil {
		return res
	}
	// 双端队列
	q := list.New()
	q.PushBack(root)
	for q.Len() > 0 {
		len := q.Len()
		var path []int
		for i := 0; i < len; i++ {
			// 移除队首第一个元素转换为TreeNode
			node := q.Remove(q.Front()).(*TreeNode)
			path = append(path, node.Val)
			if node.Left!= nil {
				q.PushBack(node.Left)
			}
			if node.Right!= nil {
				q.PushBack(node.Right)
			}
		}
		res =append(res, path)
	}
	return res
}


var ans [][]int		//全局变量

func levelOrder1(root *TreeNode) [][]int {
	if root == nil {
		//return ans   会继承上一个答案的结果
		return nil
	}
	ans =make([][]int, 0)		//每次调用函数重新初始化结果
	dfs(root, 0)
	return ans
}

func dfs(root *TreeNode, level int) {
	if root ==nil {
		return
	}
	// 计算下一层时，新建下一层
	if level > len(ans) {
		ans = append(ans,[]int{})
	}
	ans[level] =append(ans[level], root.Val)
	dfs(root.Left,level+1)
	dfs(root.Right,level+1)
}

func main() {
	
}
