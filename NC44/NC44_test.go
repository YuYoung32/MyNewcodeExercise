package NC44

import (
	"testing"
)

func Test_permute(t *testing.T) {
	cases := []struct {
		name   string
		s      string
		p      string
		expect bool
	}{
		{"normal", "aba", "*a", true},
		{"all", "aba", "*", true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := isMatch(c.s, c.p)
			if res != c.expect {
				t.Fatalf("expect %v, but %v got", c.expect, res)
			}
		})
	}
}
