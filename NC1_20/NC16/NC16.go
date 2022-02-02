package NC16

import "MyNewcodeExercise/MyTools"

type TreeNode = MyTools.TreeNode

var flag bool

func DLR(LNode, RNode *TreeNode) {
	if LNode == nil && RNode == nil {
		return
	}
	if LNode == nil || RNode == nil || LNode.Val != RNode.Val {
		flag = false
		return
	}
	DLR(LNode.Right, RNode.Left)
	DLR(LNode.Left, RNode.Right)
}

func isSymmetrical(pRoot *TreeNode) bool {
	// write code here
	flag = true
	if pRoot == nil {
		return flag
	}
	DLR(pRoot.Left, pRoot.Right)
	return flag
}
