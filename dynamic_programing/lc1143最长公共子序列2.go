package dynamic_programing

func LCS(s1 string, s2 string) string {
	// write code here
	//dp[i][j]表示 s1[:i]和s2[:j]的字符串的最长公共子串
	dp := make([][]string, len(s1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]string, len(s2)+1)
	}
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + string(s1[i-1])
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	if dp[len(s1)][len(s2)] == "" {
		return "-1"
	}
	return dp[len(s1)][len(s2)]
}
func max(a, b string) string {
	if len(a) > len(b) {
		return a
	}
	return b
}
