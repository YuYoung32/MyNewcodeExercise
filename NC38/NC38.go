package NC38

func spiralOrder(matrix [][]int) []int {
	if matrix == nil {
		return nil
	}
	var res []int
	top, left, right, btm := 0, 0, len(matrix[0])-1, len(matrix)-1
	for top < (len(matrix)+1)/2 && left < (len(matrix[0])+1)/2 {
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		for i := top + 1; i <= btm; i++ {
			res = append(res, matrix[i][right])
		}
		for i := right - 1; top != btm && i >= left; i-- {
			res = append(res, matrix[btm][i])
		}
		for i := btm - 1; left != right && i >= top+1; i-- {
			res = append(res, matrix[i][left])
		}
		left++
		right--
		btm--
		top++
	}
	return res
}
