package NC15

import (
	"MyNewcodeExercise/MyTools"
)

type TreeNode = MyTools.TreeNode

/**
 *
 * @param root TreeNode类
 * @return int整型二维数组
 */
func levelOrder(root *TreeNode) [][]int {
	// write code here
	var result [][]int
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		num := len(queue)
		tempRes := make([]int, 0)
		for i := 0; i < num; i++ {
			tempRes = append(tempRes, queue[0].Val)
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			// 去除头
			queue = queue[1:]
		}
		result = append(result, tempRes)
	}
	return result
}
