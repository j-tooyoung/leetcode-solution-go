package main

import "fmt"

// todo 未解决
// 1288. 删除被覆盖区间
// https://leetcode-cn.com/problems/remove-covered-intervals/
func removeCoveredIntervals(intervals [][]int) int {
	// [[34335,39239],[15875,91969],[29673,66453],[53548,69161],[40618,93111]]
	ans := 0
	for i := 0; i < len(intervals); i++ {
		for j := i + 1; j < len(intervals); j++ {
			if intervals[i][0] <= intervals[j][0] && intervals[i][1] >= intervals[j][1]{
				fmt.Print(i , j)
				ans++
			}else if intervals[i][0] >= intervals[j][0] && intervals[i][1] <= intervals[j][1] {
				fmt.Print(i , j)
				ans++
			}
		}
	}

	//sort.Slice(intervals, func(i, j int) bool {
	//	if intervals[i][0] == intervals[j][0] {
	//		return intervals[i][1] < intervals[j][1]
	//	}
	//	return intervals[i][0] < intervals[j][0]
	//})
	//
	return len(intervals) - ans
}

func main() {

}
