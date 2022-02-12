package NC40

import (
	"MyNewcodeExercise/MyTools"
	"testing"
)

func Test_addInList(t *testing.T) {
	cases := []struct {
		name   string
		head1  []int
		head2  []int
		expect []int
	}{
		{"normal", []int{1, 2, 3}, []int{1, 2, 3}, []int{2, 4, 6}},
		{"more count", []int{9, 9}, []int{1}, []int{1, 0, 0}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := addInList(MyTools.CreateListByArray(c.head1), MyTools.CreateListByArray(c.head2))
			if MyTools.CompareList(res, MyTools.CreateListByArray(c.expect)) == false {
				t.Fatalf("expect %v, but %v got", c.expect, res)
			}
		})
	}
}
