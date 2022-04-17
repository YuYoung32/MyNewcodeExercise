package study

import (
	"fmt"
	"testing"
)

func TestDP(t *testing.T) {
	vol := []int{1, 3, 4}
	value := []int{15, 20, 30}
	pack := 4
	fmt.Println(packProblem01OneDi(vol, value, pack))
	fmt.Println(packProblem01TwoDi(vol, value, pack))
}
