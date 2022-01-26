package NC8

import (
	"strconv"
	"testing"
)

/**
根据字符串生成树，按照节点依次生成，空节点用’#‘表示
*/
func generateTree(string2 []string) *TreeNode {
	const skip = -8288
	if string2[0][0] == '#' {
		return nil
	}
	// 将数字字符串数组转化为数字数组，再转化成节点数组
	var nodes []int
	for _, str := range string2 {
		num, err := strconv.Atoi(str)
		if err != nil {
			// 非int则做一个特殊标记，在之后生成树时给予跳过
			num = skip
		}
		nodes = append(nodes, num)
	}
	var tree []*TreeNode
	// 空占首节点。方便后续链接
	tree = append(tree, new(TreeNode))
	for _, num := range nodes {
		node := &TreeNode{
			Val:   num,
			Left:  nil,
			Right: nil,
		}
		tree = append(tree, node)
	}
	length := len(tree)
	for i := 1; i < length; i++ {
		if 2*i < length && tree[2*i].Val != skip {
			tree[i].Left = tree[2*i]
		}
		if 2*i+1 < length && tree[2*i+1].Val != skip {
			tree[i].Right = tree[2*i+1]
		}
	}
	return tree[1]
}

func compare2DArray(a1 [][]int, a2 [][]int) bool {
	compare1DArray := func(a1, a2 []int) bool {
		if len(a1) != len(a2) {
			return false
		}
		for i := 0; i < len(a1); i++ {
			if a1[i] != a2[i] {
				return false
			}
		}
		return true
	}

	if len(a1) != len(a2) {
		return false
	}

	flags := make([]bool, len(a2))
	for i := 0; i < len(flags); i++ {
		flags[i] = false
	}

	for i1 := 0; i1 < len(a1); i1++ {
		for i2 := 0; i2 < len(a2); i2++ {
			if flags[i2] == false && compare1DArray(a1[i1], a2[i2]) {
				flags[i2] = true
				break
			}
		}
	}
	for _, f := range flags {
		if f == false {
			return false
		}
	}
	return true
}

func TestFindPath(t *testing.T) {
	cases := []struct {
		Name     string
		Nodes    []string
		Number   int
		Expected [][]int
	}{

		{"not found", []string{"10", "5", "2"}, 20, [][]int{}},
		{"broaden", []string{"10"}, 10, [][]int{{10}}},
		{"normal", []string{"10", "5", "12", "4", "7"}, 22, [][]int{{10, 5, 7}, {10, 12}}},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			result := FindPath(generateTree(c.Nodes), c.Number)
			if compare2DArray(result, c.Expected) == false {
				t.Errorf("expected %v, but %v got", c.Expected, result)
			}
		})
	}
}
