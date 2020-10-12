package main

import (
	"sort"
)

// 435. 无重叠区间
// https://leetcode-cn.com/problems/non-overlapping-intervals/
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] == intervals[j][1] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})

	//fmt.Println(intervals)
	intMin := ^int(^uint(0) >> 1)
	//i, pre, ans := 0, -1*(1<<31), 0
	i, pre, ans := 0, intMin, 0
	for ;i < len(intervals);i++ {
		if pre <= intervals[i][0] {
			ans++
			pre = intervals[i][1]
		}
	}
	return len(intervals) - ans
}

func main() {
	
}
