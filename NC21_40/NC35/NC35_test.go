package NC35

import "testing"

func Test_minEditCost(t *testing.T) {
	cases := []struct {
		name   string
		str1   string
		str2   string
		ic     int
		dc     int
		rc     int
		expect int
	}{
		{"normal", "abc", "adc", 1, 2, 3, 3},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := minEditCost(c.str1, c.str2, c.ic, c.dc, c.rc)
			if res != c.expect {
				t.Fatalf("expect %v but %v got", c.expect, res)
			}
		})
	}
}
