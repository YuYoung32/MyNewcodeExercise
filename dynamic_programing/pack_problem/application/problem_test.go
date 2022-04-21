package application

import (
	"fmt"
	"testing"
	"time"
)

func TestDP(t *testing.T) {
	findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 3, 3)
}
func TestLC377(t *testing.T) {
	combinationSum4([]int{1, 2, 3}, 4)
}

func TestLC70(t *testing.T) {
	ti := time.Now()
	fmt.Println(climbStairs(44))
	fmt.Println(time.Since(ti))
	//1134903170
	//11.1255397s
}

func TestLC322(t *testing.T) {
	fmt.Println(coinChange([]int{4, 2}, 6))
}

func TestLC279(t *testing.T) {
	fmt.Println(numSquares(12))
}

func TestLC139(t *testing.T) {
	wordBreak("dogs", []string{"dog", "s", "gs"})
}
