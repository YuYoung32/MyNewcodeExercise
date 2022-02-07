package NC28

import (
	"testing"
)

func Test_minWindow(t *testing.T) {
	cases := []struct {
		name   string
		S      string
		T      string
		expect string
	}{
		{"normal", "aabbdcce", "ed", "dcce"},
		{"repeat", "aseqedeeyd", "eedd", "deeyd"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := minWindow(c.S, c.T)
			if res != c.expect {
				t.Fatalf("expect %v, but %v got", c.expect, res)
			}
		})
	}
}
