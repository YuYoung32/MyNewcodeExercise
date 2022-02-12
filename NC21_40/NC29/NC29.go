package NC29

func Find(target int, array [][]int) bool {
	//从右上角开始搜索，搜索应从每层的最大或最小开始，同时在每列的最小或最大开始，主要是确保单向流动，避免多个方向
	var x = len(array[0]) - 1
	var y = 0
	for x >= 0 && y <= len(array)-1 {
		//越界退出
		if target == array[y][x] {
			return true
		} else if target > array[y][x] {
			y++
		} else {
			x--
		}
	}
	return false
}
