package NC5

import "testing"

func generateTree(string2 string) *TreeNode {
	// 将数字字符串转化为数字数组，再转化成节点数组

	nodes := []byte(string2)
	var tree []*TreeNode
	// 空占首节点。方便后续链接
	tree = append(tree, &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	})
	for _, char := range nodes {
		node := &TreeNode{
			Val:   int(char - '0'),
			Left:  nil,
			Right: nil,
		}
		tree = append(tree, node)
	}
	length := len(tree)
	for i := 1; i < length; i++ {
		if 2*i < length {
			tree[i].Left = tree[2*i]
		}
		if 2*i+1 < length {
			tree[i].Right = tree[2*i+1]
		}
	}
	return tree[1]
}

func Test_DLR(t *testing.T) {
	cases := []struct {
		Name     string
		Nodes    string
		Expected int
	}{
		{"normal", "123", 25},
		{"long", "12034", 257},
		{"broaden", "10", 10},
		{"empty", "0", 0},
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
