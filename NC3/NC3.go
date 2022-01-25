package NC3

/**
链表中环的入口结点
推导过程：
slow一次走1，fast一次走2。
一个环形路径，一个速度快，一个速度慢，二者必定相遇，即使是离散的走。
设从开始到环入口的距离为x,从环入口到第一次相遇的距离为y,环长为n。
注意此时slow必定一圈循环都没有走过、fast必定走过至少一个循环，因为这样才能相遇。
此时slow走了x+y（不是x+y+in），fast走了x+y+in，即fast比slow多走了i圈环长
又知道fast走过长度必定是slow的两倍，则:
2(x+y)=x+y+in
得到x=in-y=(i-1)n+n-y i>=1 解释：n-y是走剩下的一点距离，再加上走完整的i-1圈的距离。
注意到slow当前位置距离环入口为y，则slow走完in-y为y+in-y=in，位置正好是在环入口位置，
注意x的定义为从开始到环入口的位置，则cur从开始走x后也到环入口位置

程序上体现出来就是cur指针从头出发，slow继续以速度1出发，走完当圈剩下的路程，当i=1时
刚好就在环入口相遇，当i>1时，slow多走了i-1圈，还是到了环入口处与slow相遇

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func EntryNodeOfLoop(pHead *ListNode) *ListNode {
	fast := pHead
	slow := pHead
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		// 没有循环节点
		return nil
	}
	cur := pHead
	for cur != slow {
		cur = cur.Next
		slow = slow.Next
	}
	return cur
}
