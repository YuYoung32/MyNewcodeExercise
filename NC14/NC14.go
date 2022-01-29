package NC14

import "MyNewcodeExercise/MyTools"

type TreeNode = MyTools.TreeNode

func reverse(list []int) []int {
	length := len(list)
	for i := 0; i <= (length-1)/2; i++ {
		list[i], list[length-1-i] = list[length-1-i], list[i]
	}
	return list
}

func Print(pRoot *TreeNode) [][]int {
	// write code here
	var result [][]int
	if pRoot == nil {
		return result
	}
	flag := 0
	queue := []*TreeNode{pRoot}
	for len(queue) != 0 {
		var tempRes []int
		num := len(queue)
		for i := 0; i < num; i++ {
			now := queue[0]
			tempRes = append(tempRes, now.Val)
			if now.Left != nil {
				queue = append(queue, now.Left)
			}
			if now.Right != nil {
				queue = append(queue, now.Right)
			}
			queue = queue[1:]
		}
		if flag%2 != 0 {
			tempRes = reverse(tempRes)
		}
		result = append(result, tempRes)
		flag++
	}
	return result
}
