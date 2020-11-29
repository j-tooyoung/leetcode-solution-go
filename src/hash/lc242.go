package main

func isAnagram(s string, t string) bool {
	if len(s) !=len(t) {
		return false
	}
	cnt := make([]int, 26)
	for i := range s {
		cnt[s[i]-'a']++
		cnt[t[i]-'a']--
	}
	for _, val := range cnt {
		if val != 0 {
			return false
		}
	}
	return true
}