package segmentArray

func countRangeSum(nums []int, lower int, upper int) int {
	length := len(nums)
	cnt := 0
	//var sum []int
	sum := make([]int, length +1)
	for i := 0; i < length; i++ {
		sum[i + 1] = sum[i] + nums[i]
	}

	for i := 0; i <= length; i++ {
		for j := 0; j < i; j++ {
			if sum[i] - sum[j] >= lower && sum[i] - sum[j] <= upper {
				cnt++
			}
		}
	}
	return cnt
}
