package NC8

import (
	"MyNewcodeExercise/MyTools"
	"testing"
)

func TestFindPath(t *testing.T) {
	cases := []struct {
		Name     string
		Nodes    []string
		Number   int
		Expected [][]int
	}{

		{"not found", []string{"10", "5", "2"}, 20, [][]int{}},
		{"broaden", []string{"10"}, 10, [][]int{{10}}},
		{"normal", []string{"10", "5", "12", "4", "7"}, 22, [][]int{{10, 5, 7}, {10, 12}}},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			result := FindPath(MyTools.GenerateTreeFromString(c.Nodes), c.Number)
			if MyTools.Compare2DArray(result, c.Expected) == false {
				t.Errorf("expected %v, but %v got", c.Expected, result)
			}
		})
	}
}
