package NC46

import (
	"sort"
)

var res [][]int
var DNum []int

func combinationSum2(num []int, target int) [][]int {
	// write code here
	sort.Ints(num)
	DNum = num
	help(0, []int{}, target)
	return res
}

func help(startIndex int, path []int, target int) {
	if target == 0 {
		res = append(res, path)
		return
	}
	for i := startIndex; i < len(DNum); i++ {
		if i > startIndex && DNum[i] == DNum[i-1] {
			//确保只有第一个可以进入运算，本级后面的重复的不计入，避免重复
			continue
		}
		if DNum[i] <= target {
			help(i+1, append(path, DNum[i]), target-DNum[i])
		} else {
			return
		}
	}
}
