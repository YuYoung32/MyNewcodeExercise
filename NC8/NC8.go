package NC8

import "MyNewcodeExercise/MyTools"

/**
切记，使用切片时，append是针对指针进行操作，已经append另一个切片后，不要再改变这个切片，否则会同步改变
例子
a:=[]int{1}
var b [][]int
append(b,a) // b为{{1}}
append(a,2) // b为{{1,2}}
*/

type TreeNode = MyTools.TreeNode

var totalPath [][]int

func DLR(root *TreeNode, path []int, expectedNumber int) {
	if root == nil {
		return
	}
	path = append(path, root.Val)
	if expectedNumber -= root.Val; expectedNumber == 0 && root.Left == nil && root.Right == nil {
		// 避免append后，指针改变污染问题
		newPath := make([]int, len(path))
		copy(newPath, path)
		totalPath = append(totalPath, newPath)
	}
	DLR(root.Left, path, expectedNumber)
	DLR(root.Right, path, expectedNumber)
}

/*FindPath
 * @param root TreeNode类
 * @param expectNumber int整型
 * @return int整型二维数组
 */
func FindPath(root *TreeNode, expectNumber int) [][]int {
	// write code here
	DLR(root, []int{}, expectNumber)
	return totalPath
}
