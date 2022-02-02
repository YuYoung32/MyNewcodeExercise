package NC1

import "testing"

func Test_solve(t *testing.T) {
	cases := []struct {
		Name     string
		S        string
		T        string
		Expected string
	}{
		{"common", "1", "2", "3"},
		{"broaden", "1", "99", "100"},
		{"normal jin", "109", "1", "110"},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if ans := solve(c.S, c.T); ans != c.Expected {
				t.Fatalf("expected %s, but %s got",
					c.Expected, ans)
			}
		})
	}
}
