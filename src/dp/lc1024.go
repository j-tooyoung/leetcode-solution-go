package main

import "math"

//
func videoStitching(clips [][]int, T int) int {
	const inf = math.MaxInt64 - 1
	dp := make([]int, T + 1)
	for i := range dp {
		dp[i] = inf
	}
	dp[0] = 0
	for i := 0; i <= T; i++ {
		for _, c := range clips {
			l, r := c[0] ,c[1]
			// 若能剪出子区间 [l,i]，则可以从 dp[l] 转移到 dp[i]
			if l < i && i <= r && dp[l] +1 < dp[i] {
				dp[i] = dp[l] + 1
			}
		}
	}
	if dp[T] == inf {
		return -1
	}
	return dp[T]
}
