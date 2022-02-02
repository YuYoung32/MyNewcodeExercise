package NC12

import (
	"MyNewcodeExercise/MyTools"
	"testing"
)

func Test_reConstructBinaryTree(t *testing.T) {
	cases := []struct {
		name          string
		pre           []int
		vin           []int
		expectedNodes []string
	}{
		{"normal", []int{1, 2, 4, 7, 3, 5, 6, 8}, []int{4, 7, 2, 1, 5, 3, 8, 6}, []string{"1", "2", "3", "4", "#", "5", "6", "#", "7", "#", "#", "#", "#", "8"}},
		{"broaden", []int{0}, []int{0}, []string{"0"}},
		{"empty", []int{}, []int{}, []string{"#"}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			flag := true
			MyTools.CompareTree(&flag, reConstructBinaryTree(c.pre, c.vin), MyTools.GenerateTreeFromString(c.expectedNodes))
			if flag == false {
				t.Fatalf("wrong tree, wrong function: reConstructBinaryTree")
			}
		})
	}
}
