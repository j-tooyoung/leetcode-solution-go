package main

import "sort"

// 逆序存放
// 法2 ： 找到非负与正数的分界线，两边扩散
func sortedSquares(A []int) []int {
	//var ans []int  报越界错误
	ans := make([]int, len(A))
	i, j, idx := 0, len(A)-1, len(A)-1
	for i <= j {
		if A[i]*A[i] < A[j]*A[j] {
			ans[idx] = A[j]*A[j]
			j--
		} else {
			ans[idx] = A[i]*A[i]
			i++
		}
		idx--
	}
	return ans
}

func sortedSquares1(A []int) []int {
	for i := 0; i < len(A); i++ {
		A[i] = A[i] * A[i]
	}
	sort.Ints(A)
	return A
}
