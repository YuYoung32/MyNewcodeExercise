package dynamic_programing

func lengthOfLIS(nums []int) int {
	//dp[i]表示，nums下标<=i时的最长序列
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = maxInt(dp[i], dp[j]+1)
			}
		}
	}
	maxdp := 1
	for _, i := range dp {
		if i > maxdp {
			maxdp = i
		}
	}
	return maxdp
}
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
