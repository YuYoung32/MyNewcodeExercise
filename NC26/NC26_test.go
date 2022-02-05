package NC26

import (
	"MyNewcodeExercise/MyTools"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	cases := []struct {
		name   string
		n      int
		expect []string
	}{
		{"normal", 2, []string{"()()", "(())"}},
		{"broaden", 1, []string{"()"}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := generateParenthesis(c.n)
			if MyTools.CompareStringArray(res, c.expect) == false {
				t.Fatalf("expect %v, but %v got", c.expect, res)
			}
		})
	}
}
