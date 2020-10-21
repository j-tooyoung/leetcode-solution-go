package main

func isLongPressedName(name string, typed string) bool {
	i, j := 0, 0
	len1, len2 := len(name), len(typed)
	for i < len1 && j < len2 {
		cnt1, cnt2 := 0, 0
		ch1,ch2 := name[i],typed[j]
		if ch1 != ch2 {
			return false
		}
		for i < len1 && name[i] == ch1  {
			i++
			cnt1++
		}
		for j < len2 && typed[j] == ch2  {
			j++
			cnt2++
		}
		if cnt2 < cnt1 {
			return false
		}
	}
	return i >= len1 && j >= len2
}
