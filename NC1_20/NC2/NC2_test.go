package NC2

import (
	"MyNewcodeExercise/MyTools"
	"testing"
)

func Test_reorderList(t *testing.T) {
	cases := []struct {
		Name     string
		Num      int
		Expected string
	}{
		{"odd", 5, "[1 5 2 4 3]"},
		{"even", 4, "[1 4 2 3]"},
		{"broaden", 1, "[1]"},
		{"long", 10, "[1 10 2 9 3 8 4 7 5 6]"},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			list := MyTools.CreateList(c.Num)
			reorderList(list)
			if MyTools.PrintList(*list) != c.Expected {
				t.Fatalf("expected %s, but %s got", c.Expected, MyTools.PrintList(*list))
			}
		})
	}
}
