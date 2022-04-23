package dynamic_programing

func findLength(nums1 []int, nums2 []int) int {
	//dp[i][j]表示 nums1中下标小于等于i 与 nums2中下标小于等于j 的最长重复序列
	dp := make([][]int, len(nums1))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(nums2))
	}
	max := 0
	//边缘初始化
	for i := 0; i < len(dp); i++ {
		if nums2[0] == nums1[i] {
			dp[i][0] = 1
			max = 1
		}
	}
	for i := 0; i < len(dp[0]); i++ {
		if nums1[0] == nums2[i] {
			dp[0][i] = 1
			max = 1
		}
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
				max = maxInt(dp[i][j], max)
			}
		}
	}
	return max
}
