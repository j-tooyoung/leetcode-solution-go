package main

func searchRange(nums []int, target int) []int {
	lo:= lowerbound(nums, target)
	hi:= upperbound(nums, target)
	return []int{lo, hi}
}

func lowerbound(nums []int, target int) int {
	return 0
}

func upperbound(nums []int, target int) int {
	return 0
}

