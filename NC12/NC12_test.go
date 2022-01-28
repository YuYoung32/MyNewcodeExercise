package NC12

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

//compareTree 比较两棵树是否相同
func compareTree(flag *bool, tree1 *TreeNode, tree2 *TreeNode) {
	if tree1 == nil && tree2 == nil {
		return
	}
	if tree1 == nil || tree2 == nil {
		*flag = false
		return
	}
	if tree1.Val != tree2.Val {
		*flag = false
		return
	}
	compareTree(flag, tree1.Left, tree2.Left)
	compareTree(flag, tree1.Right, tree2.Right)
}

func Test_reConstructBinaryTree(t *testing.T) {
	cases := []struct {
		name          string
		pre           []int
		vin           []int
		expectedNodes []string
	}{
		{"normal", []int{1, 2, 4, 7, 3, 5, 6, 8}, []int{4, 7, 2, 1, 5, 3, 8, 6}, []string{"1", "2", "3", "4", "#", "5", "6", "#", "7", "#", "#", "#", "#", "8"}},
		{"broaden", []int{0}, []int{0}, []string{"0"}},
		{"empty", []int{}, []int{}, []string{"#"}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			flag := true
			compareTree(&flag, reConstructBinaryTree(c.pre, c.vin), generateTree(c.expectedNodes))
			if flag == false {
				t.Fatalf("wrong tree, wrong function: reConstructBinaryTree")
			}
		})
	}
}
