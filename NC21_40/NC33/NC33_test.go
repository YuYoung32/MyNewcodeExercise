package NC33

import (
	"MyNewcodeExercise/MyTools"
	"testing"
)

func Test_Merge(t *testing.T) {
	cases := []struct {
		name   string
		pHead1 []int
		pHead2 []int
		expect []int
	}{
		{"normal", []int{1, 3, 5}, []int{2, 4, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"empty", []int{}, []int{1, 2, 3}, []int{1, 2, 3}},
		{"broaden", []int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := Merge(MyTools.CreateListByArray(c.pHead1), MyTools.CreateListByArray(c.pHead2))
			if MyTools.PrintList(res) != MyTools.PrintList(MyTools.CreateListByArray(c.expect)) {
				t.Fatalf("expect %v, but %v got", c.expect, res)
			}
		})
	}
}
