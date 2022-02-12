package NC23

import (
	"MyNewcodeExercise/MyTools"
	"testing"
)

func Test_partition(t *testing.T) {
	cases := []struct {
		name   string
		nodes  []int
		x      int
		expect []int
	}{
		{"normal", []int{1, 4, 3, 2, 5, 2}, 3, []int{1, 2, 2, 4, 3, 5}},
		{"broaden", []int{1}, 1, []int{1}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := partition(MyTools.CreateListByArray(c.nodes), c.x)
			if MyTools.CompareList(result, MyTools.CreateListByArray(c.expect)) == false {
				t.Fatalf("expect %v, but %v got", c.expect, MyTools.PrintList(result))
			}

		})
	}
}
