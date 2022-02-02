package NC4

import "MyNewcodeExercise/MyTools"

/**
判断链表是否有环
*/

type ListNode = MyTools.ListNode

func hasCycle(head *ListNode) bool {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
