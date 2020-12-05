package main

func leastInterval(tasks []byte, n int) int {
	if n== 0 {
		return len(tasks)
	}
	cnt := make(map[byte]int)
	for _, task := range tasks {
		cnt[task]++
	}
	maxVal := -1
	cntOfSame := 0
	for _, val := range cnt {
		val =max(val, maxVal)
	}
	for _, val := range cnt {
		if val == maxVal {
			cntOfSame++
		}
	}
	ans := cntOfSame + (n +1) *(maxVal-1)
	if ans <len(tasks) {
		return len(tasks)
	}
	return ans

}
