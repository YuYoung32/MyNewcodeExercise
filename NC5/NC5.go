package NC5

/**
二叉树求路径值之和，DLR遍历二叉树
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func DLR(root *TreeNode, pathNum int, total *int) {
	pathNum = pathNum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*total += pathNum
	}
	if root.Left != nil {
		DLR(root.Left, pathNum, total)
	}
	if root.Right != nil {
		DLR(root.Right, pathNum, total)
	}
}

func sumNumbers(root *TreeNode) int {
	// write code here
	if root == nil {
		return 0
	}
	pathNum := 0
	total := 0
	DLR(root, pathNum, &total)
	return total
}
