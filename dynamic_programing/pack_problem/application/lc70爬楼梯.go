package application

//func climbStairs(n int) int {
//	//完全背包解法
//	dp := make([]int, n+1)
//	nums := []int{1, 2}
//	dp[0] = 1
//	for i := 1; i <= n; i++ {
//		for j := 0; j < len(nums); j++ {
//			if i >= nums[j] {
//				dp[i] += dp[i-nums[j]]
//			}
//		}
//	}
//	return dp[n]
//}

//func climbStairs(n int) int {
//	//普通dp解法
//	//dp[i]表示爬到i的走法数量
//	dp := make([]int, n+1)
//	dp[0] = 1
//	dp[1] = 1
//	for i := 2; i <= n; i++ {
//		dp[i] = dp[i-1] + dp[i-2]
//	}
//	return dp[n]
//}

func climbStairs(n int) int {
	//递归暴力搜索法 超时算法
	if n == 0 {
		return 1
	}
	res := 0
	for i := 1; i <= 2; i++ {
		if n-i >= 0 {
			res += climbStairs(n - i)
		}
	}
	return res
}
