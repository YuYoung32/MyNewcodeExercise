package NC44

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	is, ip := -1, -1
	i, j := 0, 0
	for i < m {
		if j < n && (s[i] == p[j] || p[j] == '?') {
			i++
			j++
		} else if j < n && p[j] == '*' {
			is = i
			ip = j
			j++
		} else if is >= 0 {
			is++
			i = is
			j = ip + 1
		} else {
			return false
		}
	}
	for ; j < n && p[j] == '*'; j++ {

	}
	return j == n

}
