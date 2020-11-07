package array

import (
	"sort"
)

/**
https://leetcode-cn.com/problems/merge-intervals/submissions/
56. 合并区间
 */
func merge(intervals [][]int) [][]int {
	if intervals == nil || len(intervals) < 2 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	var res [][]int
	end := intervals[0][1]
	res = append(res,intervals[0])
	for i := 1; i < len(intervals); {
		if intervals[i][0] > end {
			res = append(res, intervals[i])
			end = max(intervals[i][1], end)
			i++
		} else {
			j := i
			for j < len(intervals) && intervals[j][0] <= end {
				end = max(intervals[j][1], end)
				//  fmt.Println(end)

				j++
			}
			res[len(res)-1][1] = end
			//res = append(res, []int{intervals[i-1][0], end})
			i = j
			// fmt.Println(j)
		}
	}
	return res
}

func max(i int, j int) int {
	if i < j {
		return j
	}else{
		return i
	}
}
