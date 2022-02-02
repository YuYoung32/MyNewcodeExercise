package NC7

/**
目的是求最大差
循环整个数组，先定义利润，然后每次更新最小值，然后求差，更新最大差
*/

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
 *
 * @param prices int整型一维数组
 * @return int整型
 */
func maxProfit(prices []int) int {
	// write code here
	profit := 0
	minVal := prices[0]
	for i := 1; i < len(prices); i++ {
		minVal = min(minVal, prices[i])
		profit = max(profit, prices[i]-minVal)
	}
	return profit
}
