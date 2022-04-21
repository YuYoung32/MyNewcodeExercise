package application

import (
	"math"
)

func numSquares(n int) int {
	base := int(math.Sqrt(float64(n)))
	//dp[i]表示和为i的完全平方数的最少数量
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt
		for j := 1; j <= base; j++ {
			now := j * j
			if i >= now {
				if dp[i-now] == math.MaxInt {
					continue
				}
				dp[i] = minInt(dp[i], dp[i-now]+1)
			}
		}
	}
	return dp[n]
}
