package NC19

import "testing"

func Test_FindGreatestSumOfSubArray(t *testing.T) {
	cases := []struct {
		name   string
		array  []int
		expect int
	}{
		{"normal", []int{1, -2, 3, 10, -4, 7, 2, -5}, 18},
		{"broaden", []int{1}, 1},
		{"negative", []int{-1}, 0},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := FindGreatestSumOfSubArray(c.array)
			if result != c.expect {
				t.Fatalf("ecpect %v, but %v got", c.expect, result)
			}
		})
	}
}
