package NC22

import (
	"MyNewcodeExercise/MyTools"
	"testing"
)

//切片超出len长度后不可访问，通过append可以在len长度后加入，但是会赋予新的地址，而传入的切片地址A是形参，不会改变，只能
//通过切片里的元素个数对形参指向的切片进行访问
func Test_merge(t *testing.T) {
	cases := []struct {
		name   string
		A      []int
		B      []int
		expect []int
	}{
		{"normal", []int{4, 5, 6}, []int{1, 2, 3}, []int{1, 2, 3, 4, 5, 6}},
		{"equal", []int{1, 1, 2}, []int{1, 2, 3}, []int{1, 1, 1, 2, 2, 3}},
		{"broaden", []int{1}, []int{}, []int{1}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// 创建指定容量的A切片
			temp := make([]int, len(c.A)+len(c.B), len(c.A)+len(c.B))
			for i := 0; i < len(c.A); i++ {
				temp[i] = c.A[i]
			}
			merge(temp, len(c.A), c.B, len(c.B))
			if MyTools.Compare1DArray(temp, c.expect) == false {
				t.Fatalf("expect %v, but %v got", c.expect, temp)
			}
		})
	}

}
