package NC43

import (
	"MyNewcodeExercise/MyTools"
	"testing"
)

func Test_permute(t *testing.T) {
	cases := []struct {
		name   string
		num    []int
		expect [][]int
	}{
		{"normal", []int{0, 1}, [][]int{{0, 1}, {1, 0}}},
		{"one", []int{1}, [][]int{{1}}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := permute(c.num)
			if MyTools.Compare2DArray(c.expect, res) == false {
				t.Fatalf("expect %v, but %v got", c.expect, res)
			}
		})
	}
}
