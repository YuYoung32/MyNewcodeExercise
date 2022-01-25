package NC6

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

func Test_DLR(t *testing.T) {
	cases := []struct {
		Name     string
		Nodes    []string
		Expected int
	}{
		{"normal", []string{"1", "2", "3"}, 25},
		{"long", []string{"1", "2", "0", "3", "4"}, 257},
		{"broaden", []string{"1", "0"}, 10},
		{"empty", []string{"#"}, 0},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			result := sumNumbers(generateTree(c.Nodes))
			if result != c.Expected {
				t.Fatalf("expected %v, but %v got", c.Expected, result)
			}
		})
	}
}
