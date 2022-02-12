package NC31

import "testing"

func Test_FirstNotRepeatingChar(t *testing.T) {
	cases := []struct {
		name   string
		str    string
		expect int
	}{
		{"normal", "google", 4},
		{"zero", "aa", -1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := FirstNotRepeatingChar(c.str)
			if res != c.expect {
				t.Fatalf("expect %v, but %v got", c.expect, res)
			}
		})
	}
}
