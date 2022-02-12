package NC29

import "testing"

func Test_Find(t *testing.T) {
	cases := []struct {
		name   string
		array  [][]int
		target int
		expect bool
	}{
		{"normal", [][]int{{1, 2, 8, 9}, {2, 4, 9, 12}, {4, 7, 10, 13}, {6, 8, 11, 15}}, 7, true},
		{"false", [][]int{{1, 2, 8, 9}, {2, 4, 9, 12}, {4, 7, 10, 13}, {6, 8, 11, 15}}, 66, false},
		{"broaden", [][]int{{1}}, 1, true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := Find(c.target, c.array)
			if res != c.expect {
				t.Fatalf("expect %v, but %v got", c.expect, res)
			}
		})
	}
}
