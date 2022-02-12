package NC23

import "MyNewcodeExercise/MyTools"

type ListNode = MyTools.ListNode

func partition(head *ListNode, x int) *ListNode {
	newHeadSmall := new(ListNode)
	smallH := newHeadSmall
	newHeadBig := new(ListNode)
	bigH := newHeadBig
	for head != nil {
		if head.Val >= x {
			newHeadBig.Next = head
			newHeadBig = newHeadBig.Next
		} else {
			newHeadSmall.Next = head
			newHeadSmall = newHeadSmall.Next
		}
		head = head.Next
	}
	newHeadBig.Next = nil
	newHeadSmall.Next = bigH.Next
	return smallH.Next
}
