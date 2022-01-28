package NC12

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func index(slice []int, elem int) int {
	for i, i2 := range slice {
		if i2 == elem {
			return i
		}
	}
	return -1
}

func construct(pre []int, vin []int, root *TreeNode) {
	root.Val = pre[0]
	i := index(vin, root.Val)
	if i-1 >= 0 {
		root.Left = new(TreeNode)
		construct(pre[1:1+i-1+1], vin[:i-1+1], root.Left)
	}
	if i+1 <= len(vin)-1 {
		root.Right = new(TreeNode)
		construct(pre[i+1:], vin[i+1:], root.Right)
	}
}

func reConstructBinaryTree(pre []int, vin []int) *TreeNode {
	// write code here
	if len(pre) == 0 {
		return nil
	}
	root := new(TreeNode)
	construct(pre, vin, root)
	return root
}
