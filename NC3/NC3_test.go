package NC3

import "testing"

/**
创建一个有num个节点的无头节点单链表，循环节点为第loop个
*/
func createList(num int, loop int) *ListNode {
	var head *ListNode
	head = &ListNode{
		Val:  1,
		Next: nil,
	}
	start := head
	var loopNode *ListNode = nil
	if loop == 1 {
		loopNode = head
	}
	for i := 2; i <= num; i++ {
		node := ListNode{
			Val:  i,
			Next: nil,
		}
		if i == loop {
			loopNode = &node
		}
		start.Next = &node
		start = start.Next
	}
	start.Next = loopNode
	return head
}

func Test_EntryNodeOfLoop(t *testing.T) {
	cases := []struct {
		Name string
		Num  int
		Loop int
	}{
		{"have loop", 10, 3},
		{"no loop", 2, 1},
		{"broaden", 1, -1},
		{"loop", 1, -1},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			myLoop := EntryNodeOfLoop(createList(c.Num, c.Loop))
			var myLoopValue int
			if myLoop == nil {
				myLoopValue = -1
			} else {
				myLoopValue = myLoop.Val
			}
			if myLoopValue != c.Loop {
				t.Fatalf("expected %v, but %v got", c.Loop, myLoopValue)
			}
		})
	}

}
