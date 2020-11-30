package main

import "sort"

func reorganizeString(S string) string {
	cnt := make([]int, 26)
	for _, s := range S {
		cnt[s-'a']++
	}
	var res string
	sort.Ints(cnt)
	sum := 0
	for i := 0; i +1 < len(cnt); i++ {
		sum += 0
	}
	return res
}
