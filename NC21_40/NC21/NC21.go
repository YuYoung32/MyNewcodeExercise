package NC21

import "MyNewcodeExercise/MyTools"

type ListNode = MyTools.ListNode

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	tempHead := head
	var beforeMNode *ListNode = nil
	var mNode *ListNode
	var nNode *ListNode
	var afterNNode *ListNode
	for i := 1; tempHead != nil; i++ {
		if i == m-1 {
			beforeMNode = tempHead
		}
		if i == m {
			mNode = tempHead
		}
		if i == n {
			nNode = tempHead
		}
		if i == n+1 {
			afterNNode = tempHead
			break
		}
		tempHead = tempHead.Next
	}
	// 整个过程中beforeMNode和nNode将不变
	for mNode != nNode {
		// 将mNode移到nNode与afterNNode之间
		temp := mNode.Next
		if beforeMNode != nil {
			beforeMNode.Next = mNode.Next
		} else {
			// 应对m是开头
			head = head.Next
		}
		mNode.Next = afterNNode
		nNode.Next = mNode

		afterNNode = mNode
		mNode = temp
	}
	return head
}
