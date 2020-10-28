package main


func uniqueOccurrences(arr []int) bool {
	mp := make(map[int]int)
	set := make(map[int]int)
	for _, val := range arr {
		mp[val]++
	}
	for _, val := range mp {
		if set[val] > 0 {
			return false
		}
		set[val]++
	}
	return true

}
