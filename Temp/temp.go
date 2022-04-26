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
