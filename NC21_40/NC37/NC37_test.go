package NC37

import "testing"

func CompareInterval(a []*Interval, b []*Interval) bool {
	if len(a) != len(b) {
		return false
	}
	if a == nil {
		return true
	}
	for i := 0; i < len(a); i++ {
		if a[i].Start != b[i].Start || a[i].End != b[i].End {
			return false
		}
	}
	return true
}

func Test_merge(t *testing.T) {
	cases := []struct {
		name      string
		intervals []*Interval
		expect    []*Interval
	}{
		{"normal", []*Interval{{10, 20}, {15, 30}}, []*Interval{{10, 30}}},
		{"multi", []*Interval{{10, 15}, {20, 30}}, []*Interval{{10, 15}, {20, 30}}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := merge(c.intervals)
			if CompareInterval(res, c.expect) == false {
				t.Fatalf("expect %v but %v got", c.expect, res)
			}
		})
	}
}
