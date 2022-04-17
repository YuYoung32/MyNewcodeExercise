package study

func packProblem01OneDi(vol []int, value []int, pack int) int {
	/*
		dp[i]表示背包容量为i时 背包的最大价值
		在循环遍历时，从第一个物品主键增加到最后一个物品，轮到物品j时可以理解为
		背包容量为i时，物品编号小于等于j时，背包的最大价值
	*/
	dp := make([]int, pack+1)
	for i := 0; i < len(vol); i++ {
		//反向循环防止状态重叠
		for j := pack; j >= 0; j-- {
			if j >= vol[i] {
				//要不要放入？
				dp[j] = max(dp[j], dp[j-vol[i]]+value[i])
			}
		}
	}
	return dp[pack]
}
