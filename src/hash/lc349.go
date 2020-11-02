package main

func intersection(nums1 []int, nums2 []int) []int {
	mp := make(map[int]int)
	set := make(map[int]bool)
	var res []int
	for _, val := range nums1 {
		mp[val]++
	}
	for _, val := range nums2 {
		if mp[val] > 0 && !set[val] {
			res = append(res, val)
		}
		set[val]= true
	}
	return res
}
