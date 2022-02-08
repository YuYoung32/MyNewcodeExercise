package NC32

import "testing"

func Test_sqrt(t *testing.T) {
	cases := []struct {
		name   string
		x      int
		expect int
	}{
		{"can", 4, 2},
		{"cannot", 2, 1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := sqrt(c.x)
			if res != c.expect {
				t.Fatalf("expect %v, but %v got", c.expect, res)
			}
		})
	}
}
