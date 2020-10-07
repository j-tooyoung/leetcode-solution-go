package main

import "fmt"

//
func sortColors(nums []int)  {
	
}

//一个直观的解决方案是使用计数排序的两趟扫描算法。
//首先，迭代计算出0、1 和 2 元素的个数，然后按照0、1、2的排序，重写当前数组。
func sortColors1(nums []int)  {
	var cntOf0 int
	var cntOf1 int
	for i := 0; i < len(nums); i++ {
		if nums[i]== 0 {
			cntOf0++
		} else if nums[i] == 1 {
			cntOf1++
		}
	}
	for i := 0; i < len(nums); i++ {
		if i < cntOf0 {
			nums[i] = 0
		}else if i < cntOf0 +cntOf1 {
			nums[i] = 1
		}else {
			nums[i]= 2
		}
	}
}

func main() {
	fmt.Println(13& 11)

}
