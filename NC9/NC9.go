package NC9

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var flag = false

func DLR(root *TreeNode, sum int) {
	if root == nil {
		return
	}
	sum -= root.Val
	if sum == 0 && root.Left == nil && root.Right == nil {
		flag = true
		return
	}
	DLR(root.Left, sum)
	DLR(root.Right, sum)
}
func hasPathSum(root *TreeNode, sum int) bool {
	DLR(root, sum)
	return flag
}
