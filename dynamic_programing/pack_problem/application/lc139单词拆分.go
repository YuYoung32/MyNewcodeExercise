package application

func wordBreak(s string, wordDict []string) bool {
	//dp[i]表示 s[:i+1]内是否能分割
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := i; j > 0; j-- {
			for _, str := range wordDict {
				if s[j-1:i] == str {
					dp[i] = dp[j-1] || dp[i]
				}
			}
		}
	}
	return dp[len(s)]
}
