package NC27

import "sort"

var result [][]int
var path []int
var length int
var array []int

func subsets(A []int) [][]int {
	// 避免help传入过多不变的参数
	sort.Ints(A)
	length = len(A)
	array = A
	help(0)
	return result
}

func help(start int) {
	if start >= length {
		newPath := make([]int, len(path))
		copy(newPath, path)
		result = append(result, newPath)
		return
	}
	//不添加
	help(start + 1)
	//添加一个
	path = append(path, array[start])
	help(start + 1)
	path = path[:len(path)-1]
}
