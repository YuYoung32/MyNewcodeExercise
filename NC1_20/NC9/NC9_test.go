package NC9

import (
	"MyNewcodeExercise/MyTools"
	"testing"
)

func Test_hasPathSum(t *testing.T) {
	cases := []struct {
		name   string
		input  []string
		sum    int
		expect bool
	}{
		{"normal", []string{"1", "2", "0"}, 1, true},
		{"broaden", []string{"1"}, 1, true},
		{"empty", []string{"#"}, 0, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			root := MyTools.GenerateTreeFromString(c.input)
			res := hasPathSum(root, c.sum)
			if res != c.expect {
				t.Errorf("expect %t,but got %t", c.expect, res)
			}
		})
	}
}
