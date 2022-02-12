package NC38

import (
	"MyNewcodeExercise/MyTools"
	"testing"
)

func Test_spiralOrder(t *testing.T) {
	cases := []struct {
		name   string
		matrix [][]int
		expect []int
	}{
		{"normal", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{"small", [][]int{{1, 2}, {3, 4}}, []int{1, 2, 4, 3}},
		{"big", [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}, []int{1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10}},
		{"broaden", [][]int{{2, 3}}, []int{2, 3}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := spiralOrder(c.matrix)
			if MyTools.Compare1DArray(res, c.expect) == false {
				t.Fatalf("expect %v, but %v got", c.expect, res)
			}
		})
	}
}
