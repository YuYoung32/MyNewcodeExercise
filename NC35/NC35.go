package NC35

func minEditCost(str1 string, str2 string, ic int, dc int, rc int) int {
	// write code here
	m := make([][]int, len(str1)+1)
	for i := 0; i < len(str1)+1; i++ {
		m[i] = make([]int, len(str2)+1)
		m[i][0] = i * dc
	}
	for i := 0; i < len(str2)+1; i++ {
		m[0][i] = i * ic
	}
	for i := 1; i < len(str1)+1; i++ {
		for j := 1; j < len(str2)+1; j++ {
			if str1[i-1] == str2[j-1] {
				m[i][j] = m[i-1][j-1]
			} else {
				m[i][j] = min(min(m[i-1][j]+dc, m[i][j-1]+ic), m[i-1][j-1]+rc)
			}
		}
	}
	return m[len(str1)][len(str2)]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
