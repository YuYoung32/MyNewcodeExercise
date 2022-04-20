package dynamic_programing

func maxProfit(prices []int) int {
	//dp[i][0]表示在第i天 不持有股票(0) 手中资金数(可以为负)
	dp := make([][]int, len(prices))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][1] = -prices[0]
	for i := 1; i < len(dp); i++ {
		//昨日无股票 昨日有股票但今日卖出
		dp[i][0] = maxInt(dp[i-1][0], dp[i-1][1]+prices[i])
		//昨日无股票今日买入股票 昨日有股票今日先卖后买
		dp[i][1] = maxInt(dp[i-1][0]-prices[i], dp[i-1][1])
	}
	return maxInt(dp[len(dp)-1][0], dp[len(dp)-1][1])
}
