package dynamic_programing

func numDistinct(s string, t string) int {
	//dp[i][j]表示s[:i]中包含t[:j]的个数
	dp := make([][]int, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(t)+1)
		//每个s必定包含一个空串“”
		dp[i][0] = 1
	}

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			if s[i-1] == t[j-1] {
				//选择当前，或者不选择当前
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				//不选择当前
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(s)][len(t)]
}
