package NC39

var res int
var col = make(map[int]int, 10)
var leftSlash = make(map[int]int, 10)
var rightSlash = make(map[int]int, 10)
var num int

func Nqueen(n int) int {
	// write code here
	num = n
	help(0)
	return res
}

func help(line int) {
	if line == num {
		res++
	}
	for i := 0; i < num; i++ {
		_, okc := col[i]
		_, okl := leftSlash[line-i]
		_, okr := rightSlash[line+i]
		if okc == false && okl == false && okr == false {
			col[i] = 1
			leftSlash[line-i] = 1
			rightSlash[line+i] = 1
			help(line + 1)
			delete(col, i)
			delete(leftSlash, line-i)
			delete(rightSlash, line+i)
		}
	}
}
