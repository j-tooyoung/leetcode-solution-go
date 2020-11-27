package sort

import (
	"math"
	"sort"
)

// 164. 最大间距
func maximumGap(nums []int) int {
	if nums == nil || len(nums) <= 1 {
		return 0
	}
	sort.Ints(nums)
	ans:= math.MinInt32
	for i := 1; i < len(nums); i++ {
		ans = max(ans, nums[i] - nums[i-1])
	}
	return ans
}

func max(i int, j int) int {
	if i < j {
		return j
	}
	return i
}
