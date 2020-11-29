package array

import "sort"

func largestPerimeter(A []int) int {
	sort.Ints(A)
	for i := len(A) - 1; i >= 0; i-- {
		j := i - 2
		if j >= 0 && A[i -1] + A[j] > A[i] && A[i] - A[i-1] < A[j] && A[i] - A[j] < A[i-1] {
			return A[i] + A[i -1] + A[j]
		}
	}
	return 0
}
