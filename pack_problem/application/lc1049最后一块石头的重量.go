package application

func lastStoneWeightII(stones []int) int {
	//目的是分出两堆质量尽量相同的石头，然后相减即可，同上一题
	sum := 0
	for _, stone := range stones {
		sum += stone
	}
	pack := sum / 2
	dp := make([]int, pack+1)
	for i := 0; i < len(stones); i++ {
		for j := pack; j >= 0; j-- {
			if j >= stones[i] {
				dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
			}
		}
	}
	return sum - 2*dp[pack]
}
