package dynamic_programing

func countSubstrings(s string) int {
	//dp[i][j]表示s[i:j+1]内的回文子串个数
	dp := make([][]bool, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s)+1)
	}
	cnt := 0
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				//单个或两个
				if j-i <= 1 {
					cnt++
					dp[i][j] = true
				} else {
					//看内部是不是回文串
					if dp[i+1][j-1] {
						cnt++
						dp[i][j] = true
					}
				}
			}
		}
	}
	return cnt
}

/*传统扩展解法
func countSubstrings(s string) int {
    res:=0
    for i:=0;i<len(s);i++{
        res+=extend(s,i,i)
        res+=extend(s,i,i+1)
    }
    return res
}

func extend(s string,i,j int)int{
    res:=0
    for i>=0&&j<len(s){
        if s[i]==s[j]{
            res++
            i--
            j++
        }else{
            break
        }
    }
    return res
}

*/
