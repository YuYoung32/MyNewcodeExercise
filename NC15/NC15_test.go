package NC15

import (
	"MyNewcodeExercise/MyTools"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	cases := []struct {
		name   string
		nodes  []string
		expect [][]int
	}{
		{"normal", []string{"1", "2", "3"}, [][]int{{1}, {2, 3}}},
		{"broaden", []string{"1"}, [][]int{{1}}},
		{"empty", []string{"#"}, [][]int{}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := levelOrder(MyTools.GenerateTreeFromString(c.nodes))
			if MyTools.StrictCompare2DArray(result, c.expect) == false {
				t.Fatalf("expect %v, but %v got", c.expect, result)
			}
		})
	}
}
