package dynamic_programing

import "fmt"

func longestPalindromeSubseq(s string) int {
	//dp[i][j]表示s[i:j+1]中的最长回文串
	dp := make([][]int, len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(s))
	}
	for i := len(dp) - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < len(dp[0]); j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				//单独加入一个s[i]或s[j]
				dp[i][j] = maxInt(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	for _, ints := range dp {
		fmt.Println(ints)
	}
	return dp[0][len(s)-1]
}
