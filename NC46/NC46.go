package NC46

import "sort"

func combinationSum2(num []int, target int) [][]int {
	// write code here
	sort.Ints(num)
	n := len(num)
	result := [][]int{}
	visited := make([]bool, n)
	var tb func(temp []int, start, t int)
	tb = func(temp []int, start, t int) {
		if t == target {
			tempCopy := make([]int, len(temp))
			copy(tempCopy, temp)
			result = append(result, tempCopy)
			return
		}
		if start >= n || t > target {
			return
		}
		for i := start; i < n; i++ {
			if i > 0 && num[i] == num[i-1] && !visited[i-1] {
				continue
			}
			temp = append(temp, num[i])
			visited[i] = true
			tb(temp, i+1, t+num[i])
			visited[i] = false
			temp = temp[:len(temp)-1]
		}
	}
	tb([]int{}, 0, 0)
	return result
}
