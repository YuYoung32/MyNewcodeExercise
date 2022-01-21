package NC2

/**
隔差逆转单链表
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *
 * @param head ListNode类
 * @return void
 */
func reorderList(head *ListNode) {
	// write code here
	// 快慢指针确认中点
	fast := head
	slow := head
	if fast == nil || fast.Next == nil {
		return
	}
	for fast.Next != nil {
		fast = fast.Next
		if fast.Next == nil {
			break
		}
		fast = fast.Next
		slow = slow.Next
	}
	// 颠倒后段链表
	end := fast         // 记录旧链表的结束即新链表开始的位置
	newEnd := slow.Next // 记录新链表的结束位置
	for newEnd.Next != nil {
		temp := newEnd.Next
		newEnd.Next = newEnd.Next.Next

		temp.Next = slow.Next
		slow.Next = temp
	}
	// 把后段插入到前面
	start := head
	for end != nil {
		temp := end.Next

		slow.Next = slow.Next.Next

		end.Next = start.Next
		start.Next = end

		start = start.Next.Next
		end = temp
	}
}
