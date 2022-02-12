package NC24

import (
	"MyNewcodeExercise/MyTools"
	"testing"
)

func Test_partition(t *testing.T) {
	cases := []struct {
		name   string
		nodes  []int
		expect []int
	}{
		{"normal", []int{1, 2, 2, 3, 4, 4, 5}, []int{1, 3, 5}},
		{"nil", []int{1, 1}, []int{}},
		{"continuous", []int{2, 2, 2, 3, 3, 4}, []int{4}},
		{"broaden", []int{1}, []int{1}},
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
