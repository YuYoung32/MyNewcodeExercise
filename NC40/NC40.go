package NC40

import "MyNewcodeExercise/MyTools"

type ListNode = MyTools.ListNode

func addInList(head1 *ListNode, head2 *ListNode) *ListNode {
	head1, h1l := reverse(head1)
	head2, h2l := reverse(head2)
	//结果存在head1当中，确保长度head1长
	if h1l < h2l {
		head1, head2 = head2, head1
	}
	ret := head1
	jin := 0
	ben := 0
	for head1 != nil && head2 != nil {
		ben = (head1.Val + head2.Val + jin) % 10
		jin = (head1.Val + head2.Val + jin) / 10
		head1.Val = ben
		head1 = head1.Next
		head2 = head2.Next
	}
	for head1 != nil {
		ben = (head1.Val + jin) % 10
		jin = (head1.Val + jin) / 10
		head1.Val = ben
		head1 = head1.Next
	}
	if jin != 0 {
		temp := ret
		//定位到最后一个，因为nil指针统一指向一个地址，不能更改，需要手动定位
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = &ListNode{
			Val:  jin,
			Next: nil,
		}
	}
	ret, _ = reverse(ret)
	return ret
}

func reverse(head *ListNode) (*ListNode, int) {
	if head == nil {
		return nil, 0
	}
	before := &ListNode{
		Val:  0,
		Next: head,
	}
	after := head.Next
	length := 1
	for after != nil {
		temp := after.Next
		head.Next = head.Next.Next
		after.Next = before.Next
		before.Next = after
		after = temp
		length++
	}
	return before.Next, length
}
