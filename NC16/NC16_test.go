package NC16

import (
	"MyNewcodeExercise/MyTools"
	"testing"
)

func Test_isSymmetrical(t *testing.T) {
	cases := []struct {
		name   string
		nodes  []string
		expect bool
	}{
		{"norma false", []string{"1", "2", "3"}, false},
		{"normal", []string{"1", "2", "2", "3", "4", "4", "3"}, true},
		{"broaden", []string{"1"}, true},
		{"empty", []string{"#"}, true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := isSymmetrical(MyTools.GenerateTreeFromString(c.nodes))
			if result != c.expect {
				t.Fatalf("expect %v, but %v got", c.expect, result)
			}
		})
	}
}
