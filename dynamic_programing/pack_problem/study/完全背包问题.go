package study

import "fmt"

//LC518 零钱兑换 https://leetcode-cn.com/problems/coin-change-2/
func change(amount int, coins []int) int {
	//dp[i]表示 能够组成i的组合数
	dp := make([]int, amount+1)
	dp[0] = 1
	//先容量后物品是求排列数 因为会有重合
	//先物品后容量是组合数 因为物品只会有一轮机会 不重合
	for i := 0; i < len(coins); i++ {
		for j := 1; j <= amount; j++ {
			if j >= coins[i] {
				//为什么不是相乘？
				//这里不是对前一个状态的相加而是不同物品的相加，维度不同
				dp[j] += dp[j-coins[i]]
			}
		}
	}
	fmt.Println(dp)
	return 0
}

/*
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
*/
