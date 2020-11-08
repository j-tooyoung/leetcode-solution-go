package main

//122. 买卖股票的最佳时机 II
func maxProfit_ii(prices []int) int {
	//var dp [len(prices)][2]int  报错
	dp := make([][2]int, len(prices))
	//dp := new([][]int)
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0],dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1],dp[i-1][0]-prices[i])
	}
	return dp[len(prices)-1][0]
}