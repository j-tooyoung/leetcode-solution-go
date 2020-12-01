package main

func reorganizeString(S string) string {
	cnt := make([]int, 26)
	maxCnt := -1
	for _, s := range S {
		cnt[s-'a']++
	}
	idx:= 0
	for i, val := range cnt {
		if maxCnt < val {
			idx = i
			maxCnt = val
		}
	}
	if maxCnt > len(S) >>1 {
		return ""
	}
	res := make([]byte,len(S))
	for i := 1; i < len(S); i+= 2 {
		for cnt[i] >0 {
			cnt[i]--
			res[idx] = byte('a' + i)
			idx+= 2
		}
	}
	for i := 1; i < len(S); i++ {
	}
	return string(res)
}

func max(i int, j int) int {
	if i < j {
		return j
	}
	return i
}
