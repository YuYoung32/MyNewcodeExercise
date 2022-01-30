package NC17

import "testing"

func Test_getLongestPalindrome(t *testing.T) {
	cases := []struct {
		name    string
		quesStr string
		expect  int
	}{
		{"odd", "ababa", 5},
		{"even", "abbba", 5},
		{"long", "abbcbbdd", 5},
		{"broaden", "a", 1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := getLongestPalindrome(c.quesStr)
			if result != c.expect {
				t.Fatalf("expect %v, but %v got", c.expect, result)
			}
		})
	}
}
