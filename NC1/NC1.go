package NC1

/**
以字符串的形式读入两个数字，编写一个函数计算它们的和，以字符串形式返回。

数据范围：len(s),len(t) \le 100000len(s),len(t)≤100000，字符串仅由'0'~‘9’构成
要求：时间复杂度 O(n)O(n)
示例1
输入：
"1","99"
返回值：
"100"
说明：
1+99=100

示例2
输入：
"114514",""
返回值：
"114514"
*/

func solve(s string, t string) string {
	// write code here
	// 确保长的是s
	sb := []byte(s)
	tb := []byte(t)
	sbl := len(sb)
	tbl := len(tb)
	if sbl < tbl {
		sb, tb = tb, sb
		sbl, tbl = tbl, sbl
	}
	i := sbl - 1
	j := tbl - 1
	jin := 0
	ben := 0
	for i >= 0 {
		if j >= 0 {
			ben = (int(sb[i]-'0') + int(tb[j]-'0') + jin) % 10
			jin = (int(sb[i]-'0') + int(tb[j]-'0') + jin) / 10
		} else {
			ben = (int(sb[i]-'0') + jin) % 10
			jin = (int(sb[i]-'0') + jin) / 10
		}
		sb[i] = byte(ben) + '0'
		i--
		j--
	}
	if jin > 0 {
		return "1" + string(sb)
	}
	return string(sb)
}
