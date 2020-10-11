package main

func maxProfit(prices []int) int {
	//ans ,len:=0,len(prices)- 1
	ans:=0
	var nums []int
	for i := 0; i < len(prices)-1; i++ {
		nums[i] = prices[i + 1] -prices[i]
	}
	//for i := 0; i < len(nums); i++ {
	//	if nums[i] > 0 {
	//		ans += nums[i]
	//	}
	//}
	return ans
}

func main() {
	
}
