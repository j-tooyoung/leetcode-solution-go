package main


func partitionLabels(S string) []int {
	var res []int
	mp := make(map[uint8]int)
	for i := 0; i < len(S); i++ {
		mp[S[i]] = i
	}
	start, end := 0, 0
	for i := 0; i < len(S); {
		end = max(end, mp[S[i]])
		if i == end {
			res = append(res, end -start + 1)
			start=end + 1
		}
	}
	return res
}

func partitionLabels1(S string) []int {
	var res []int
	mp := make(map[uint8]int)
	for i := 0; i < len(S); i++ {
		mp[S[i]] = i
	}
	cur := 0
	for i := 0; i < len(S); {
		cur = mp[S[i]]
		j, last := i, cur
		// 不断获取区间的最大值，直到区间所有值的最后位置均在此区间内
		for j < last {
			cur = max(cur, mp[S[j]])
			if cur > last {
				last = cur
			}
			j++
		}
		res = append(res, cur-i+1)
		i = cur + 1
	}
	return res
}

func max(i int, j int) int {
	if i < j {
		return j
	} else {
		return i
	}
}
