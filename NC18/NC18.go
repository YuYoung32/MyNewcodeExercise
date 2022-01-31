package NC18

func rotateMatrix(mat [][]int, n int) [][]int {
	// 先根据左对角线对称交换，再左右对称交换
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			mat[i][j], mat[j][i] = mat[j][i], mat[i][j]
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			mat[i][j], mat[i][n-j-1] = mat[i][n-j-1], mat[i][j]
		}
	}
	return mat
}
