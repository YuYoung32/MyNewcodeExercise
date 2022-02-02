package NC6

import (
	"MyNewcodeExercise/MyTools"
	"testing"
)

func Test_maxPathSum(t *testing.T) {
	cases := []struct {
		Name     string
		Nodes    []string
		Expected int
	}{
		{"normal", []string{"1", "2", "3"}, 6},
		{"negative numbers", []string{"-2", "#", "-3", "-5"}, -2},
		{"long", []string{"-20", "8", "20", "#", "#", "15", "6"}, 41},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			result := maxPathSum(MyTools.GenerateTreeFromString(c.Nodes))
			if result != c.Expected {
				t.Fatalf("expected %v, but %v got", c.Expected, result)
			}
		})
	}
}
