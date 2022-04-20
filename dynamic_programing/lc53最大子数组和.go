package dynamic_programing

func maxSubArray(nums []int) int {
	//dp[i]表示nums下标<=i时的最大子数组的和
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(dp); i++ {
		dp[i] = maxInt(nums[i], nums[i]+dp[i-1])
	}
	res := int(-10e5)
	for _, i := range dp {
		if i > res {
			res = i
		}
	}
	return res
}
