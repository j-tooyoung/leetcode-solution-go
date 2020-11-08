package array

import (
	"sort"
)

func countBits(n int) int {
	count := 0
	for n != 0 {
		count += n & 1
		n >>= 1
	}
	return count
}

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		a, b := countBits(arr[i]), countBits(arr[j])
		if a == b {
			return arr[i] < arr[j]
		}
		return a < b
	})
	return arr
}

//作者：xiao_ben_zhu
//链接：https://leetcode-cn.com/problems/sort-integers-by-the-number-of-1-bits/solution/fu-xi-wei-yun-suan-fu-1356-gen-ju-shu-zi-er-jin-zh/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func sortByBits1(arr []int) []int {
	cntOf1 := make(map[int]int)
	for _, val := range arr {
		cntOf1[val] = getBitCnt(val)		// 越界
	}
	sort.Slice(arr,func(i, j int) bool {
		if cntOf1[arr[i]] == cntOf1[arr[j]] {
			return arr[i] < arr[j]
		}
		return cntOf1[arr[i]] < cntOf1[arr[j]]
	})
	return arr
}

func getBitCnt(val int) int {
	cnt := 0
	for val > 0 {
		val = val & (val - 1)
		cnt++
	}
	return cnt
}

