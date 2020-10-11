package main

import (
	"sort"
	"strconv"
)

//var dp []bool
////var dp []int

// dfs 时间TLE
//func canPartition(nums []int) bool {
//	sum := 0
//	// 遍历的索引下标
//	//for i := range nums {
//	for i := 0; i < len(nums); i++ {
//		sum += nums[i]
//	}
//	//sort.Ints(nums)
//	//fmt.Println(sum)
//	if sum % 2 != 0 {
//		return false
//	}
//	sum = sum >>1
//	var dp [len(nums)][sum + 1]int
//
//	return dp[len(nums)][sum]
//
//}
var mp map[string]bool

func canPartition1(nums []int) bool {

	mp = make(map[string]bool)
	sum := 0
	// 遍历的索引下标
	//for i := range nums {
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	//sort.Ints(nums)
	//fmt.Println(sum)
	if sum%2 != 0 {
		return false
	}
	if sum == 0 {
		return true
	}
	sum >>= 1
	//sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	sort.Ints(nums)
	return dfs(nums, 0, sum)

}

func dfs(nums []int, idx int, sum int) bool {
	if sum ==0 {
		return true
	}
	if sum < 0 || idx >= len(nums) {
		return false
	}
	if nums[idx] < sum {
		return false
	}
	key := strconv.Itoa( idx) + "," + strconv.Itoa(sum)
	i, ok := mp[key]
	if ok {
		return i
	}
	if sum == 0 {
		return true
	}
	flag := dfs(nums, idx+1, sum-nums[idx]) || dfs(nums, idx+1, sum)
	mp[key] = flag
	return flag
}

func main() {

	var nums = []int{100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,99,97}
	canPartition1(nums)
}
