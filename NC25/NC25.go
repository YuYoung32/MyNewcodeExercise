package NC25

import "MyNewcodeExercise/MyTools"

type ListNode = MyTools.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	retHead := head
	for head != nil && head.Next != nil {
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}
	return retHead
}
