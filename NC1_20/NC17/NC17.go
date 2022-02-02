package NC17

// 中心扩散法求最长回文
func getLongestPalindrome(A string) int {
	// write code here
	if len(A) == 0 {
		return 0
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	result := 0
	for i := 0; i < len(A); i++ {
		l := i - 1
		r := i + 1
		//确认中心，考虑整个回文为偶数个和奇数个情况，若为偶数个中心必然为n个重复，若为奇数个则直接赋值
		//step是指为了跳出这个回文中心，把它看作是下一个回文的一部分，
		//step可有可无，因为是在回文中心下次判断得出同样结果，如果是中心的最后一个（右边缘）则必不再是回文，因此可以直接跳过。
		//删除step一样代码可以通过
		for r < len(A) && A[i] == A[r] {
			r++
		}
		for l >= 0 && r < len(A) && A[l] == A[r] {
			l--
			r++
		}
		result = max(result, r-l-1)
	}
	return result
}
