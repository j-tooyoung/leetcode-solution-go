package main

import (
	"fmt"
	"sort"
)

// 452. 用最少数量的箭引爆气球
//https://leetcode-cn.com/problems/minimum-number-of-arrows-to-burst-balloons/
// 合并重复区间，计算个数
func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	for _, point := range points {
		fmt.Print(point[0],point[1])
	}
	ans := 1
	length := len(points)
	if length < 1 {
		return 0
	}
	end := points[0][1]
	// 用例 [[2,3],[2,3]] 有问题 ？
	//[[1,2],[4,5],[1,5]]
	//if length == 2 {
	//	if points[1][0] <= points[0][1] {
	//		return ans
	//	}else{
	//		return ans+ 1
	//	}
	//}
	for i := 1; i < length; {
		if points[i][0] <= end {
			for i <length && points[i][0] <= end {
				i++
			}
			ans++
		}else{
			ans++
			end = points[i][1]
			i++
		}
	}
	return ans
}

func main() {

}
