package NC20

import "strconv"

var result []string
var paths []int

func restoreIpAddresses(s string) []string {
	// write code here
	find(s, 0)
	return result
}

func find(s string, startIndex int) {
	//每轮都尝试从start开始，1 2 3位依次去尝试
	for i := startIndex; i < startIndex+3 && i < len(s); i++ {
		if len(paths) == 3 {
			// 全部匹配，最后一个必须是剩下的全部
			if num, _ := strconv.Atoi(s[startIndex:]); (s[startIndex] != '0' || s[startIndex:] == "0") && num >= 0 && num <= 255 {
				paths = append(paths, num)
				str := strconv.Itoa(paths[0]) + "." + strconv.Itoa(paths[1]) + "." + strconv.Itoa(paths[2]) + "." + strconv.Itoa(paths[3])
				result = append(result, str)
				paths = paths[:len(paths)-1]
			}
			return
		}
		if num, _ := strconv.Atoi(s[startIndex : i+1]); (s[startIndex] != '0' || s[startIndex:i+1] == "0") && num >= 0 && num <= 255 {
			// 符合条件则添加，进入下一段IP的判断
			paths = append(paths, num)
			find(s, i+1)
			// 本次判断后恢复现场，方便下次判断
			paths = paths[:len(paths)-1]
		} else {
			// 不符合条件直接进入下一段IP的判断，因为越来越大必不符合条件，此else return可加可不加
			return
		}
	}
}
