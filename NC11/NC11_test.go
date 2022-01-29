package NC11

import (
	"MyNewcodeExercise/MyTools"
	"testing"
)

func Test_sortedArrayToBSD(t *testing.T) {
	cases := []struct {
		name string
		num  []int
	}{
		{"normal", []int{1, 2, 3}},
		{"broaden", []int{1}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if MyTools.CheckIsBST(sortedArrayToBST(c.num)) == false {
				t.Fatalf("wrong result, bad function: sortedArrayToBST")
			}
		})
	}
}
