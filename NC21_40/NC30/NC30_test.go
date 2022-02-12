package NC30

import (
	"testing"
)

func Test_minNumberDisappeared(t *testing.T) {
	cases := []struct {
		name   string
		array  []int
		expect int
	}{
		{"normal", []int{3, 2, 1}, 4},
		{"zero", []int{0}, 1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := minNumberDisappeared(c.array)
			if res != c.expect {
				t.Fatalf("expect %v, but %v got", c.expect, res)
			}
		})
	}
}
