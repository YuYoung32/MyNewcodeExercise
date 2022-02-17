package NC46

import (
	"MyNewcodeExercise/MyTools"
	"testing"
)

func Test_combinationSum2(t *testing.T) {
	cases := []struct {
		name   string
		num    []int
		target int
		expect [][]int
	}{
		{"normal", []int{1, 1, 2}, 2, [][]int{{1, 1}, {2}}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := combinationSum2(c.num, c.target)
			if MyTools.StrictCompare2DArray(res, c.expect) == false {
				t.Fatalf("expect %v, but %v got", c.expect, res)
			}
		})
	}
}
