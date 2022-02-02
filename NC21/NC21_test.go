package NC21

import (
	"MyNewcodeExercise/MyTools"
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	cases := []struct {
		name        string
		nodes       []int
		m           int
		n           int
		expectNodes []int
	}{
		{"normal", []int{1, 2, 3, 4, 5}, 2, 4, []int{1, 4, 3, 2, 5}},
		{"broaden", []int{1}, 1, 1, []int{1}},
		{"total reverse", []int{1, 2, 3, 4, 5}, 1, 5, []int{5, 4, 3, 2, 1}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := reverseBetween(MyTools.CreateListByArray(c.nodes), c.m, c.n)
			expect := MyTools.CreateListByArray(c.expectNodes)
			if MyTools.CompareList(result, expect) == false {
				t.Fatalf("expect %v, but %v got", MyTools.PrintList(expect), MyTools.PrintList(result))
			}
		})
	}
}
