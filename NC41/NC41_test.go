package NC41

import "testing"

func Test_maxLength(t *testing.T) {
	cases := []struct {
		name   string
		arr    []int
		expect int
	}{
		{"normal", []int{1, 2, 3, 4, 1, 5, 6, 7, 8, 1}, 8},
		{"zero", []int{}, 0},
		{"repeat", []int{1, 1, 1}, 1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := maxLength(c.arr)
			if res != c.expect {
				t.Fatalf("expect %v, but %v got", c.expect, res)
			}
		})
	}
}
