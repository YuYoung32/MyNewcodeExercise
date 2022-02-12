package NC36

import "testing"

func Test_findMedianinTwoSortedAray(t *testing.T) {
	cases := []struct {
		name   string
		arr1   []int
		arr2   []int
		expect int
	}{
		{"normal", []int{1, 2, 3}, []int{4, 5, 6}, 3},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := findMedianinTwoSortedAray(c.arr1, c.arr2)
			if res != c.expect {
				t.Fatalf("expect %v but %v got", c.expect, res)
			}
		})
	}
}
