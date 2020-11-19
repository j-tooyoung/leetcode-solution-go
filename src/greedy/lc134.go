package main

// https://leetcode-cn.com/problems/gas-station/
// 134. 加油站
// greedy
func canCompleteCircuit(gas []int, cost []int) int {
	return -1
}

// 暴力模拟
func canCompleteCircuit1(gas []int, cost []int) int {
	i := 0
	for ; i < len(gas); i++ {
		cnt := gas[i] - cost[i]
		j := (i + 1) % len(gas)
		for cnt >= 0{
			cnt += gas[j] - cost[j]
			if cnt < 0 {
				break;
			}
			j = (j + 1) % len(gas)
			if j == i {
				return i
			}
		}
	}
	return -1
}
