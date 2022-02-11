package NC39

import (
	"testing"
)

func Test_Nqueen(t *testing.T) {
	cases := []struct {
		name   string
		n      int
		expect int
	}{
		{"broaden", 1, 1},
		{"normal", 9, 352},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := Nqueen(c.n)
			if res != c.expect {
				t.Fatalf("expect %v, but %v got", c.expect, res)
			}
		})
	}
}
