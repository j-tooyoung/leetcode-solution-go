package main

import "sort"

// todo improve perfermence
//左指针左边均为非零数；
//	右指针左边直到左指针处均为零。
func moveZeroes(nums []int) {
	i, j := 0, 0
	for j < len(nums) {
		if nums[j] != 0 {
			nums[j],nums[i] = nums[i],nums[j]
			j++
			i++
		}else{
			j++
		}
	}
}

func moveZeroes1(nums []int) {
	i, j := 0, 0
	for ; i < len(nums); {
		if nums[i] == 0 {
			j := i
			for j < len(nums) && nums[j] == 0 {
				j++
			}
			if j < len(nums) {
				nums[j], nums[i] = nums[i], nums[j]
			}
			i++
		} else {
			i++
			j++
		}
	}
	return
}

// 暴力解法 fixme error
func moveZeroes2(nums []int) {
	sort.Slice(nums, func(i, j int) bool {
		a, b := nums[i], nums[j]
		if a == 0 && b != 0 {
			return false
		}
		if b == 0 && a != 0 {
			return true
		}
		if a!= 0 && b != 0 {
			return i < j
		}
		return false
	})
}
