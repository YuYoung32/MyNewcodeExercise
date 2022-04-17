package application

import "fmt"

func findMaxForm(strs []string, n int, m int) int {
	//dp[i][j]表示 最多具有i个0和j个1的 子集个数 dp[m][n]最多m个1最多n个0的子集个数
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < len(strs); i++ {
		os, zs := getOsZs(strs[i])
		for j := m; j >= os; j-- {
			for k := n; k >= zs; k-- {
				dp[j][k] = max(dp[j][k], dp[j-os][k-zs]+1)
			}
		}
	}
	for _, ints := range dp {
		fmt.Println(ints)
	}
	return dp[m][n]
}

func getOsZs(str string) (os, zs int) {
	for _, c := range str {
		if c == '1' {
			os++
		} else if c == '0' {
			zs++
		}
	}
	return
}
