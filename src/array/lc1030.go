package array

import "sort"

func allCellsDistOrder(n int, m int, r0 int, c0 int) [][]int {
	ans := make([][]int, 0, n*m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			ans = append(ans, []int{i, j})
		}
	}
	sort.Slice(ans, func(i, j int) bool {
		a, b := ans[i], ans[j]
		return abs(a[0]-r0)+abs(a[1]-c0) < abs(b[0]-r0)+abs(b[1]-c0)
	})
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}


