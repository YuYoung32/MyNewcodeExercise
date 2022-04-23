package dynamic_programing

func findLengthOfLCIS(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	max := 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
			max = maxInt(dp[i], max)
		}
	}
	return max
}
