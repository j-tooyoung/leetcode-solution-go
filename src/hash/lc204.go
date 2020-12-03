package main

//
func countPrimes(n int) int {
	//var mp map[int]int
	mp := make(map[int]int)

	cnt:= 0
	for i := 2; i < n; i++ {
		if mp[i] == 0 {
			cnt++
			// 2i,3i,4i,5i
			for j := i+ i; j < n; j+= i {
				mp[j] = 1
			}
		}
	}
	return cnt

}
