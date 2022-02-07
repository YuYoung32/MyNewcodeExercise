package NC28

func minWindow(S string, T string) string {
	//匹配成功的节点
	type TP struct {
		index int
		pair  []int
	}
	var path []TP
	for i := 0; i < len(S); i++ {
		//判断当前字符是否在T内
		temp := TP{
			index: i,
			pair:  nil,
		}
		for j := 0; j < len(T); j++ {
			if S[i] == T[j] {
				temp.pair = append(temp.pair, j)
			}
		}
		if temp.pair != nil {
			path = append(path, temp)
		}
	}
	if len(path) < len(T) {
		return ""
	}
	var retStrs []string
	//找到所有匹配的字符串
	for start := 0; start < len(path); start++ {
		flags := make([]bool, len(T))
		for i := start; i < len(path); i++ {
			for _, pairIndex := range path[i].pair {
				if flags[pairIndex] == false {
					flags[pairIndex] = true
					break
				}
			}
			if check(flags) == true {
				retStrs = append(retStrs, S[path[start].index:path[i].index+1])
				break
			}

		}
	}
	minStr := retStrs[0]
	for _, str := range retStrs {
		if len(str) < len(minStr) {
			minStr = str
		}
	}
	return minStr
}
func check(flags []bool) bool {
	for _, flag := range flags {
		if flag == false {
			return false
		}
	}
	return true
}
