package dynamic_programing

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	//dp[i]表示偷 下标小于等于i的房价所获得的最大值
	bbefore := nums[0]
	before := maxInt(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		before, bbefore = maxInt(before, bbefore+nums[i]), before
	}
	return before
}
