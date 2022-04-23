package dynamic_programing

func isSubsequence(s string, t string) bool {
	//本题转化为求s和t的最长公共子序列，如果最长为s的长度 则返回true
	dp := make([][]int, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(t)+1)
	}
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				//不相等则继承上一个 相当于是“删除”
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	if dp[len(s)][len(t)] == len(s) {
		return true
	}
	return false
}
