package NC32

func sqrt(x int) int {
	for i := 0; i <= x; i++ {
		if i*i == x {
			return i
		}
		if i*i > x {
			return i - 1
		}
	}
	return 0
}
