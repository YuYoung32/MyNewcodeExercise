package NC13

import "MyNewcodeExercise/MyTools"

type TreeNode = MyTools.TreeNode

var maxd = 0

func DLR(root *TreeNode, depth int) {
	if root == nil {
		return
	}
	depth = depth + 1
	if depth > maxd {
		maxd = depth
	}
	DLR(root.Left, depth)
	DLR(root.Right, depth)
}

func maxDepth(root *TreeNode) int {
	// write code here
	depth := 0
	DLR(root, depth)
	return maxd
}
