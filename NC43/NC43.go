package NC43

var dNum []int
var res [][]int
var path []int
var total int
var count []int

func permute(num []int) [][]int {
	// write code here
	total = len(num)
	dNum = num
	count = make([]int, len(dNum))
	help()
	return res
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
