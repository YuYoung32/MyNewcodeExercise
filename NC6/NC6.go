package NC6

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var total = -1000

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func maxPath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 左右子树的最大路径值
	l := maxPath(root.Left)
	r := maxPath(root.Right)
	// 记录最大路径，如果有更大的则替换掉，注意当左右为负数的时候直接记录当前节点值，不要再加上l或r
	sum := root.Val
	if l > 0 {
		sum += l
	}
	if r > 0 {
		sum += r
	}
	total = max(total, sum)
	// 向上级提供，左中或右中或中，两个max的作用是l或r是负数时只返回当前节点
	return max(root.Val, max(r, l)+root.Val)
}

/**
 *
 * @param root TreeNode类
 * @return int整型
 */
func maxPathSum(root *TreeNode) int {
	// write code here
	if root == nil {
		return total
	}
	maxPath(root)
	return total
}
