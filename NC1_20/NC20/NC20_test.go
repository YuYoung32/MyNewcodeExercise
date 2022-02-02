package NC20

import (
	"MyNewcodeExercise/MyTools"
	"testing"
)

func Test_restoreIpAddresses(t *testing.T) {
	cases := []struct {
		name   string
		inputS string
		expect []string
	}{
		{"normal", "25525522135", []string{"255.255.22.135", "255.255.221.35"}},
		{"single", "1111", []string{"1.1.1.1"}},
		{"before 0", "000256", []string{}},
		{"many before 0", "010010", []string{"0.10.0.10", "0.100.1.0"}},
		{"special", "0000", []string{"0.0.0.0"}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := restoreIpAddresses(c.inputS)
			if MyTools.CompareStringArray(result, c.expect) == false {
				t.Fatalf("expect %v, but %v got", c.expect, result)
			}
		})
	}
}
