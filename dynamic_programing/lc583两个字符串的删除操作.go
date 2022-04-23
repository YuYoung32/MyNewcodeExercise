package dynamic_programing

func minDistance(word1 string, word2 string) int {
	//dp[i][j]表示word1[:i]到word[:j]的最小步数
	dp := make([][]int, len(word1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(word2)+1)
		dp[i][0] = i
	}
	for i := 0; i < len(dp[0]); i++ {
		dp[0][i] = i
	}

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				//选择两个中的一个进行删除，或者两个都删除
				dp[i][j] = minInt(dp[i-1][j]+1, minInt(dp[i][j-1]+1, dp[i-1][j-1]+2))
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
