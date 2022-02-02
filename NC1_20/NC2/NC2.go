package NC2

import "MyNewcodeExercise/MyTools"

type ListNode = MyTools.ListNode

/**
 *
 * @param head ListNode类
 * @return void
 */
func reorderList(head *ListNode) {
	// write code here
	fast := head
	slow := head
	// 确认至少有一个节点
	if head == nil || head.Next == nil {
		return
	}
	// 快慢指针确认中点
	for fast.Next != nil {
		fast = fast.Next
		if fast.Next == nil {
			break
		}
		fast = fast.Next
		slow = slow.Next
	}
	// 颠倒后段链表
	end := fast         // 记录旧链表的结束即新链表开始的位置，用于下一步插入到前面
	newEnd := slow.Next // 记录新链表的结束位置
	for newEnd.Next != nil {
		// 提取出后个
		temp := newEnd.Next
		newEnd.Next = newEnd.Next.Next
		// 插入到前面，slow的下一个，slow不用移动
		temp.Next = slow.Next
		slow.Next = temp
	}
	// 把后段插入到前面
	start := head
	for end != nil {
		// 临时存放下一个需要取出的节点
		temp := end.Next
		// 取出颠倒的链表的首个节点-从原链表移除-插入前面链表
		slow.Next = slow.Next.Next
		end.Next = start.Next
		start.Next = end
		// 确定下一个需要插入的位置
		start = start.Next.Next

		end = temp
	}
}
