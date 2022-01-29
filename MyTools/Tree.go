package MyTools

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//GenerateTreeFromString 根据字符串生成树，按照节点依次生成，空节点用’#‘表示
func GenerateTreeFromString(string2 []string) *TreeNode {
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

func ChildNumber(root *TreeNode) int {
	num := 0
	if root.Left != nil {
		num++
	}
	if root.Right != nil {
		num++
	}
	return num
}

func CheckIsSorted(root *TreeNode) bool {
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
	num := ChildNumber(root)
	if num == 0 {
		return
	}
	if num == 1 && lastNum == 1 {
		// 不再平衡
		*flag = false
		return
	}
	if CheckIsSorted(root) == false {
		// 不再有序
		*flag = false
	}
	DLR(root.Left, flag, num)
	DLR(root.Right, flag, num)
}

//CheckIsBST 判断是否是平衡搜索二叉树，要求root.Left.Val <= root.Val <= root.Right.Val
func CheckIsBST(root *TreeNode) bool {
	flag := true
	//初始化时假设root有两个孩子，避免root只有一个孩子的情况被误判为false
	child := 2
	DLR(root, &flag, child)
	return flag
}

//CompareTree 比较两棵树是否相同
func CompareTree(flag *bool, tree1 *TreeNode, tree2 *TreeNode) {
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
	CompareTree(flag, tree1.Left, tree2.Left)
	CompareTree(flag, tree1.Right, tree2.Right)
}
