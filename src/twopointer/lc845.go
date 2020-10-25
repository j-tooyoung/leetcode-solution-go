package main

func longestMountain(A []int) int {
	res := 0
	for index, val := range A {
		i, j := index-1, index+1
		for i >= 0 && A[i] < val {
			val = A[i]
			i--
		}
		val = A[index]
		for j < len(A) && A[j] < val {
			val = A[j]
			j++
		}
		if i == index-1 || j == index+1 {
			res = max(res, 0)
		} else {
			res = max(res, j-i-1)
		}
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
