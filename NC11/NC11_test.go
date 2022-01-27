package NC11

import (
	"testing"
)

func childNumber(root *TreeNode) int {
	num := 0
	if root.Left != nil {
		num++
	}
	if root.Right != nil {
		num++
	}
	return num
}

func checkIsSorted(root *TreeNode) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}
	if root.Left == nil && root.Right != nil {
		if root.Right.Val < root.Val {
			return false
		} else {
			return true
		}
	}
	if root.Left != nil && root.Right == nil {
		if root.Left.Val > root.Val {
			return false
		} else {
			return true
		}
	}
	if root.Left.Val <= root.Val && root.Right.Val >= root.Val {
		return true
	} else {
		return false
	}
}

func DLR(root *TreeNode, flag *bool, lastNum int) {
	num := childNumber(root)
	if num == 0 {
		return
	}
	if num == 1 && lastNum == 1 {
		// 不再平衡
		*flag = false
		return
	}
	if checkIsSorted(root) == false {
		// 不再有序
		*flag = false
	}
	DLR(root.Left, flag, num)
	DLR(root.Right, flag, num)
}

//checkIsBSD 判断是否是平衡搜索二叉树，要求root.Left.Val <= root.Val <= root.Right.Val
func checkIsBSD(root *TreeNode) bool {
	flag := true
	//初始化时假设root有两个孩子，避免root只有一个孩子的情况被误判为false
	child := 2
	DLR(root, &flag, child)
	return flag
}

func Test_sortedArrayToBSD(t *testing.T) {
	cases := []struct {
		name string
		num  []int
	}{
		{"normal", []int{1, 2, 3}},
		{"broaden", []int{1}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if checkIsBSD(sortedArrayToBST(c.num)) {
				t.Fatalf("wrong result, bad function: checkIsBSD")
			}
		})
	}
}
