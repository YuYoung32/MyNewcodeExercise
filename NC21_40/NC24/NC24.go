package NC24

import "MyNewcodeExercise/MyTools"

type ListNode = MyTools.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	// write code here
	if head == nil {
		return nil
	}
	trueHead := head
	retHead := head
	start := head
	end := head
	beforeStart := &ListNode{
		Val:  0,
		Next: head,
	}
	for head.Next != nil {
		if head.Next.Val == head.Val {
			// 相同则同步记录，之后一并删除
			end = end.Next
		} else {
			if start != end {
				// 有重复，可删除
				if beforeStart.Next == trueHead {
					// 头部被删除则后移
					trueHead = end.Next
					retHead = end.Next
				}
				// 删除节点链
				beforeStart.Next = end.Next
				// 新一轮判断
				start = end.Next
				end = end.Next
			} else {
				// 无重复，不用删除，直接开始新一轮判断
				beforeStart = beforeStart.Next
				start = start.Next
				end = end.Next
			}
		}
		head = head.Next
	}
	// 重复到最后一个节点的情况
	if start != end {
		// 有重复，可删除
		if beforeStart.Next == trueHead {
			// 头部被删除则后移
			retHead = end.Next
		}
		// 删除节点链
		beforeStart.Next = end.Next
	}
	return retHead
}
