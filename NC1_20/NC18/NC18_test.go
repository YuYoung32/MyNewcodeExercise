package NC18

import (
	"MyNewcodeExercise/MyTools"
	"testing"
)

func Test_rotateMatrix(t *testing.T) {
	cases := []struct {
		name      string
		matrix    [][]int
		dimension int
		expect    [][]int
	}{
		{"normal", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 3, [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}},
		{"broaden", [][]int{{1}}, 1, [][]int{{1}}},
		{"empty", [][]int{}, 0, [][]int{}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if MyTools.Compare2DArray(rotateMatrix(c.matrix, c.dimension), c.expect) == false {
				t.Fatalf("wrong answer, bad function: totateMatrix")
			}
		})
	}
}
