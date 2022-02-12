package NC27

import (
	"MyNewcodeExercise/MyTools"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	cases := []struct {
		name   string
		A      []int
		expect [][]int
	}{
		{"normal", []int{1, 2}, [][]int{{}, {1}, {2}, {1, 2}}},
		{"broaden", []int{1}, [][]int{{}, {1}}},
		{"empty", []int{}, [][]int{{}}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := subsets(c.A)
			if MyTools.Compare2DArray(c.expect, res) == false {
				t.Fatalf("expect %v, but %v got", c.expect, res)
			}
		})
	}
}
