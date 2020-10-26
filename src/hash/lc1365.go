package main

func smallerNumbersThanCurrent(nums []int) []int {
	//res := make([]int,len(nums))
	//cnt := make([]int, 110)
	//sum := make([]int, 110)
	//for i := 0; i < len(nums); i++ {
	//	cnt[nums[i]]++
	//}
	//
	//for i := 1; i <= 100; i++ {
	//	cnt[i] += cnt[i - 1]
	//	sum[i] = cnt[i -1]
	//}
	//
	//for i := 0; i < len(nums); i++ {
	//	res[i] = sum[nums[i]]
	//}
	//return res
	res := make([]int,len(nums))
	cnt := make([]int, 110)
	for i := 0; i < len(nums); i++ {
		cnt[nums[i]]++
	}

	for i := 1; i <= 100; i++ {
		cnt[i] += cnt[i - 1]
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			res[i] = 0
		}else{
			res[i] = cnt[nums[i] - 1]
		}
	}
	return res
}

func smallerNumbersThanCurrent1(nums []int) []int {
	res := make([]int,len(nums))

	for i := 0; i < len(nums); i++ {
		cnt := 0
		for j := 0; j < len(nums); j++ {
			if j != i && nums[j] < nums[i] {
				cnt++
			}
		}
		res[i] = cnt
	}
	return res
}

func main() {
	nums := []int{3, 5, 4, 1}
	smallerNumbersThanCurrent(nums)
}
