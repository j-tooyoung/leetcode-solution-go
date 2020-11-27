package main

import "fmt"

// 四數相加ii
func fourSumCount(A []int, B []int, C []int, D []int) int {
	res := 0
	m1 := make(map[int]int,510)
	m2 := make(map[int]int,510)
	for _, v1 := range A {
		for _, v2 := range B {
			m1[v1+v2]++
		}
	}
	for _, v1 := range C {
		for _, v2 := range D {
			m2[v1+v2]++
		}
	}

	for index := range m1 {
		fmt.Println(index,m2[-index])

		res += m1[index] * m2[-index]
	}
	return res
}
