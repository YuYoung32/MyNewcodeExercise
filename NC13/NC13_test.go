package NC13

import (
	"MyNewcodeExercise/MyTools"
	"testing"
)

func Test_maxDepth(t *testing.T) {
	cases := []struct {
		name     string
		nodes    []string
		expected int
	}{
		{"normal", []string{"1", "2", "3"}, 2},
		{"broaden", []string{"1"}, 1},
		{"empty", []string{"#"}, 0},
		{"long", []string{"1", "2", "3", "4", "#", "#", "5"}, 3},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := maxDepth(MyTools.GenerateTreeFromString(c.nodes))
			if result != c.expected {
				t.Fatalf("expected %v, but %v got", c.expected, result)
			}
		})
	}
}
