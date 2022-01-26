package NC7

import "testing"

func Test_maxProfit(t *testing.T) {
	cases := []struct {
		Name     string
		Prices   []int
		Expected int
	}{
		{"normal", []int{2, 4, 1}, 2},
		{"broaden", []int{1, 2}, 1},
		{"no profit", []int{3, 2, 1}, 0},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if result := maxProfit(c.Prices); result != c.Expected {
				t.Fatalf("expected %v, but %v got", c.Expected, result)
			}
		})
	}
}
