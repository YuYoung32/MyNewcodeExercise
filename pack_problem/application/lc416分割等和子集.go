package application

/*
华为笔试20220407 第三题
给你一个只包含正整数的非空数组nums 。
请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
*/
func canPartition(nums []int) bool {
	//若设背包容积为sum(nums)/2，每项的价值和体积都是nums[i]，那么有且只有取得最大价值的同时又刚好填满
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}
	pack := sum / 2
	//dp[i]表示容积为i时的最大价值
	dp := make([]int, pack+1)
	for i := 0; i < len(nums); i++ {
		for j := pack; j >= 0; j-- {
			if j >= nums[i] {
				dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
			}
		}
	}
	if dp[pack] == pack {
		return true
	}
	return false
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
