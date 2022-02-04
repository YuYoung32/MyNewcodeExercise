package NC25

import (
	"MyNewcodeExercise/MyTools"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	cases := []struct {
		name   string
		nodes  []int
		expect []int
	}{
		{"normal", []int{1, 2, 2, 3}, []int{1, 2, 3}},
		{"broaden", []int{1, 1, 1}, []int{1}},
		{"empty", []int{}, []int{}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := deleteDuplicates(MyTools.CreateListByArray(c.nodes))
			if MyTools.CompareList(result, MyTools.CreateListByArray(c.expect)) == false {
				t.Fatalf("expect %v, but %v got", c.expect, MyTools.PrintList(result))
			}
		})
	}
}
