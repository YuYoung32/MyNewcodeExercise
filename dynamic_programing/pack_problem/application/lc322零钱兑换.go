package application

import (
	"math"
)

func coinChange(coins []int, amount int) int {
	//dp[i]表示要凑够i 所需要的最少硬币个数
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] {
				//说明不是由首个推导而来，不可以凑整
				if dp[i-coins[j]] == math.MaxInt {
					continue
				}
				dp[i] = minInt(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt {
		return -1
	}
	return dp[amount]
}
func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
