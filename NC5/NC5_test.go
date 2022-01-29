package NC5

import (
	"MyNewcodeExercise/MyTools"
	"testing"
)

func Test_DLR(t *testing.T) {
	cases := []struct {
		Name     string
		Nodes    []string
		Expected int
	}{
		{"normal", []string{"1", "2", "3"}, 25},
		{"long", []string{"1", "2", "0", "3", "4"}, 257},
		{"broaden", []string{"1", "0"}, 10},
		{"empty", []string{"0"}, 0},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			result := sumNumbers(MyTools.GenerateTreeFromString(c.Nodes))
			if result != c.Expected {
				t.Fatalf("expected %v, but %v got", c.Expected, result)
			}
		})
	}
}
