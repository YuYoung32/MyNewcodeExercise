package NC26

//递归模板：终止条件，计入总体条件，递归变量改变与递归函数
var result []string
var total int

func help(path []byte, lNum int, rNum int) {
	if lNum > total || rNum > total || rNum > lNum {
		return
	}
	if len(path) == 2*total {
		result = append(result, string(path))
	}
	help(append(path, '('), lNum+1, rNum)
	help(append(path, ')'), lNum, rNum+1)
}

func generateParenthesis(n int) []string {
	total = n
	help([]byte{}, 0, 0)
	return result
}
