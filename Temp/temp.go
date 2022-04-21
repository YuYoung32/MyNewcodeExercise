package main

import "fmt"

func change(amount int, coins []int) int {
	//dp[i]表示 能够组成i的组合数
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] {
				dp[i] = dp[i] + dp[i-coins[j]]
			}
		}
	}
	fmt.Println(dp)
	return 0
}
