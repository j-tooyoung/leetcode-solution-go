package main

import "fmt"

// 75. 颜色分类
// https://leetcode-cn.com/problems/sort-colors/
//链接：https://leetcode-cn.com/problems/sort-colors/solution/kuai-su-pai-xu-partition-guo-cheng-she-ji-xun-huan/
func sortColors(nums []int)  {
	if len(nums) < 2 {
		return
	}
	// all in [0, zero) = 0
	// all in [zero, i) = 1
	// all in [two, len - 1] = 2

	// 循环终止条件是 i == two，那么循环可以继续的条件是 i < two
	// 为了保证初始化的时候 [0, zero) 为空，设置 zero = 0，
	// 所以下面遍历到 0 的时候，先交换，再加
	zero,two,i := 0,len(nums), 0
	for i < two {
		if nums[i] == 0 {
			nums[zero], nums[i] = nums[i], nums[zero]
			//fmt.Println(zero, i, nums[zero],nums[i])
			// why todo [zero,i)全为1
			i++
			zero++
		}else if nums[i] == 1 {
			i++
		}else {
			two--
			nums[two], nums[i] = nums[i],nums[two]
			fmt.Println(two, i, nums[two],nums[i])

		}
	}

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
