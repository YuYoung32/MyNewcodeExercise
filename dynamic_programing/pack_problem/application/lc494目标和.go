package application

func findTargetSumWays(nums []int, target int) int {
	/*可以分解为全部数相加，正数部分和负数部分，
	根据方程 sump-sumn=target,sump+sumn=total可以知道需要凑成的正数 sump=(target+sum)/2
	*/
	total := 0
	for _, num := range nums {
		total += num
	}
	S := total + target
	if S%2 == 1 || S < 0 {
		return 0
	}
	psum := S / 2
	//dp[i]表示目标结果为i时，合成目标target的种类数
	dp := make([]int, psum+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := psum; j >= 0; j-- {
			if j >= nums[i] {
				dp[j] += dp[j-nums[i]]
			}
		}
	}
	return dp[psum]
}
