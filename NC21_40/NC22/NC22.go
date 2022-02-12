package NC22

func merge(A []int, m int, B []int, n int) {
	var res []int
	i := 0
	j := 0
	for i < m && j < n {
		if A[i] <= B[j] {
			res = append(res, A[i])
			i++
		} else {
			res = append(res, B[j])
			j++
		}
	}
	for i < m {
		res = append(res, A[i])
		i++
	}
	for j < n {
		res = append(res, B[j])
		j++
	}
	for i = 0; i < len(res); i++ {
		A[i] = res[i]
	}
}
