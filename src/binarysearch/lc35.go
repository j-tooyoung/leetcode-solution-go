package main

import "fmt"

func searchInsert(nums []int, target int) int {
	i, j := 0, len(nums) -1
	for i < j {
		mid := (i + j + 1) >> 1
		if nums[mid] <= target {
			i = mid
		} else {
			j = mid - 1
		}
	}
	fmt.Println(nums[j])

	if nums[i] > target {
		return i+1
	}else {
		return i
	}

}

func main() {

}
