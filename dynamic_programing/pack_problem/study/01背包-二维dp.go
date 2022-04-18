package study

func packProblem01TwoDi(vol []int, value []int, pack int) int {
	//dp[i][j]表示 编号小于等于i的物品 放在 大小为pack的背包里 产生的最大价值
	dp := make([][]int, len(vol))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, pack+1)
	}
	//第0行，只放第一个物品时的情况 第一纵行是背包大小为0的情况，都为0
	for i := 0; i <= pack; i++ {
		if i >= vol[0] {
			dp[0][i] = value[0]
		}
	}
	//遍历
	for i := 1; i < len(dp); i++ {
		for j := 1; j <= pack; j++ {
			if j >= vol[i] {
				//要不要把当前物品加入背包？不加入直接等于上级，加入则等于上级减少相应空间加入当前物品价值
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-vol[i]]+value[i])
			} else {
				dp[i][j] = dp[i-1][j]
			}

		}
	}
	return dp[len(dp)-1][pack]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
