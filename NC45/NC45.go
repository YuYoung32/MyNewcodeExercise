package NC45

import "MyNewcodeExercise/MyTools"

type TreeNode = MyTools.TreeNode

var path1, path2, path3 []int

func threeOrders(root *TreeNode) [][]int {
	// write code here
	DLR(root)
	LDR(root)
	LRD(root)
	return [][]int{path1, path2, path3}
}

func DLR(root *TreeNode) {
	if root == nil {
		return
	}
	path1 = append(path1, root.Val)
	DLR(root.Left)
	DLR(root.Right)
}

func LDR(root *TreeNode) {
	if root == nil {
		return
	}
	LDR(root.Left)
	path2 = append(path2, root.Val)
	LDR(root.Right)
}

func LRD(root *TreeNode) {
	if root == nil {
		return
	}
	LRD(root.Left)
	LRD(root.Right)
	path3 = append(path3, root.Val)
}
