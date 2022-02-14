package NC42

import "sort"

var dNum []int
var res [][]int
var path []int
var total int
var count []int

func permuteUnique(num []int) [][]int {
	// write code here
	sort.Ints(num)
	total = len(num)
	dNum = num
	count = make([]int, len(dNum))
	help()
	return deRepeat()
}

func help() {
	if len(path) == total {
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)
		return
	}
	for i := 0; i < total; i++ {
		if count[i] == 0 {
			path = append(path, dNum[i])
			count[i] = 1
			help()
			count[i] = 0
			path = path[:len(path)-1]
		}
	}
}

func deRepeat() [][]int {
	var newRes [][]int
	m := make(map[string]int)
	for _, r := range res {
		str := ""
		for _, n := range r {
			str = str + string(byte(n))
		}
		if m[str] == 0 {
			newRes = append(newRes, r)
			m[str] = 1
		}
	}
	return newRes
}
