package NC19

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func FindGreatestSumOfSubArray(array []int) int {
	// write code here
	sum := -1000
	now := 0
	for i := 0; i < len(array); i++ {
		now += array[i]
		sum = max(sum, now)
		if now < 0 {
			now = 0
		}
	}
	return sum
}
