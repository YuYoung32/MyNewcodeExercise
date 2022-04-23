package dynamic_programing

import "testing"

func TestLC53(t *testing.T) {
	maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
}

func TestLC718(t *testing.T) {
	findLength([]int{1, 2, 3, 2, 8},
		[]int{5, 6, 1, 4, 7})
}

func TestLC115(t *testing.T) {
	numDistinct("baegg", "bag")
}

func TestLC516(t *testing.T) {
	longestPalindromeSubseq("bbbab")
}
