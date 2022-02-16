package NC45

import (
	"MyNewcodeExercise/MyTools"
	"testing"
)

func Test_threeOrders(t *testing.T) {
	cases := []struct {
		name   string
		root   []string
		expect [][]int
	}{
		{"normal", []string{"1", "2", "3"}, [][]int{{1, 2, 3}, {2, 1, 3}, {2, 3, 1}}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := threeOrders(MyTools.GenerateTreeFromString(c.root))
			if MyTools.StrictCompare2DArray(res, c.expect) == false {
				t.Fatalf("expect %v, but %v got", c.expect, res)
			}
		})
	}
}
