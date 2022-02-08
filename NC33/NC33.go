package NC33

import "MyNewcodeExercise/MyTools"

type ListNode = MyTools.ListNode

func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	newHead := new(ListNode)
	pointer := newHead
	for pHead1 != nil && pHead2 != nil {
		if pHead1.Val <= pHead2.Val {
			pointer.Next = pHead1
			pointer = pointer.Next
			pHead1 = pHead1.Next
		} else {
			pointer.Next = pHead2
			pointer = pointer.Next
			pHead2 = pHead2.Next
		}
	}
	for pHead1 != nil {
		pointer.Next = pHead1
		pointer = pointer.Next
		pHead1 = pHead1.Next
	}
	for pHead2 != nil {
		pointer.Next = pHead2
		pointer = pointer.Next
		pHead2 = pHead2.Next
	}
	return newHead.Next
}
