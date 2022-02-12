package NC34

import "testing"

func Test_uniquePaths(t *testing.T) {
	cases := []struct {
		name   string
		m      int
		n      int
		expect int
	}{
		{"normal", 2, 2, 2},
		{"broaden", 1, 1, 1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := uniquePaths(c.m, c.n)
			if res != c.expect {
				t.Fatalf("expect %v but %v got", c.expect, res)
			}
		})
	}
}
