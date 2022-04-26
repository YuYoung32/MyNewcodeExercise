package dynamic_programing

func longestValidParentheses(s string) int {
	//dp[i]表示s[:i+1]的最长合法括号数量
	dp := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		//()和))
		if s[i] == ')' {
			//上一个是)
			if s[i-1] == '(' {
				if i >= 2 {
					//直接从本组前一个加2
					dp[i] = dp[i-2] + 2
				} else {
					//前方无，相当于0+2
					dp[i] = 2
				}
				//或者上一组的前一个是(
			} else if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' { //跳转到上一组完整的前一个字符
				//从该组前一个加2 再加上内部嵌入的组
				if i-dp[i-1]-2 >= 0 {
					dp[i] = dp[i-dp[i-1]-2] + 2 + dp[i-1]
				} else {
					//前方无，相当于0+2
					dp[i] = dp[i-1] + 2
				}

			}
		}
	}
	max := 0
	for _, r := range dp {
		if max < r {
			max = r
		}
	}
	return max
}

/*一种简单的解法
package main

func longestValidParentheses(s string) int {
	res, l, r := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			l++
		} else {
			r++
		}
		if l == r {
			res = max(res, l*2)
		} else if r > l {
			l, r = 0, 0
		}
	}
	l, r = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			l++
		} else {
			r++
		}
		if l == r {
			res = max(res, l*2)
		} else if l > r {
			l, r = 0, 0
		}
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}



*/
