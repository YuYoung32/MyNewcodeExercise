package NC41

func maxLength(arr []int) int {
	max := 0
	m := map[int]int{}
	temp := 0
	for i := 0; i < len(arr); i++ {
		index, ok := m[arr[i]]
		temp++
		//记录索引
		m[arr[i]] = i
		if ok {
			if temp-1 > max {
				max = temp - 1
			}
			temp = 0
			i = index
			m = map[int]int{}
		}
	}
	if temp > max {
		max = temp
	}
	return max
}
