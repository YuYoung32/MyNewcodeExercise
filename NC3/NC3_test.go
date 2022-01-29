package NC3

import (
	"MyNewcodeExercise/MyTools"
	"testing"
)

func Test_EntryNodeOfLoop(t *testing.T) {
	cases := []struct {
		Name string
		Num  int
		Loop int
	}{
		{"have loop", 10, 3},
		{"no loop", 2, 1},
		{"broaden", 1, -1},
		{"loop", 1, -1},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			myLoop := EntryNodeOfLoop(MyTools.CreateLoopList(c.Num, c.Loop))
			var myLoopValue int
			if myLoop == nil {
				myLoopValue = -1
			} else {
				myLoopValue = myLoop.Val
			}
			if myLoopValue != c.Loop {
				t.Fatalf("expected %v, but %v got", c.Loop, myLoopValue)
			}
		})
	}

}
