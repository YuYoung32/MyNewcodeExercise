package NC11

import (
	"MyNewcodeExercise/MyTools"
)

type TreeNode = MyTools.TreeNode

func generate(root *TreeNode, nodes []int) {
	mid := len(nodes) / 2
	root = &TreeNode{
		Val:   nodes[mid],
		Left:  nil,
		Right: nil}
	if mid-1 > 0 {
		generate(root.Left, nodes[:mid-1])
	}
	if mid+1 <= len(nodes)-1 {
		generate(root.Right, nodes[mid+1:])
	}
}

func sortedArrayToBST(num []int) *TreeNode {
	// write code here
	if len(num) == 0 {
		return nil
	}
	root := new(TreeNode)
	generate(root, num)
	return root
}
