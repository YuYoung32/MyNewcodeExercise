package NC30

func minNumberDisappeared(nums []int) int {
	m := make([]int, 5e5)
	for _, num := range nums {
		if num > 0 {
			m[num]++
		}
	}
	m[0] = 1
	var ret int
	for index, num := range m {
		if num == 0 {
			ret = index
			break
		}
	}
	return ret
}
