package NC4

import (
	"MyNewcodeExercise/MyTools"
	"testing"
)

func Test_hasCycle(t *testing.T) {
	cases := []struct {
		Name     string
		Num      int
		Loop     int
		Expected bool
	}{
		{"have loop", 10, 3, true},
		{"no loop", 2, -1, false},
		{"broaden no loop", 1, -1, false},
		{"broaden have loop", 1, 1, true},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			flag := hasCycle(MyTools.CreateLoopList(c.Num, c.Loop))
			if flag != c.Expected {
				t.Fatalf("expected %v, but %v got", c.Expected, flag)
			}
		})
	}

}
