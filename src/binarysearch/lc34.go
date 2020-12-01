package main

import "sort"

func searchRange1(nums []int, target int) []int {
	lo := sort.SearchInts(nums, target)
	if lo == len(nums) || nums[lo] != target {
		return []int{-1, -1}
	}
	hi := sort.SearchInts(nums, target+1) - 1
	return []int{lo, hi}
}

func searchRange(nums []int, target int) []int {
	if nums == nil || len(nums) < 1{
		return []int{-1,-1}
	}
	lo := lowerbound(nums, target)
	// fmt.Println(lo)
	hi := upperbound(nums, target)
	return []int{lo, hi}
	// lo := sort.SearchInts(nums, target)
	// if lo == len(nums) || nums[lo] != target {
	// 	return []int{-1, -1}
	// }
	// hi := sort.SearchInts(nums, target+1) - 1
	// return []int{lo, hi}
}

//第一个>=target的最小数
func lowerbound(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := (lo + hi) /2
		if nums[mid] >= target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	if nums[lo] != target {
		return -1
	}
	return lo
}

func upperbound(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := (lo + hi + 1) >> 1
		if nums[mid] <= target {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	if nums[lo] != target {
		return -1
	}
	return lo
}
