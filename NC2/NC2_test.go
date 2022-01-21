package NC2

import (
	"fmt"
	"testing"
)

func createList(num int) *ListNode {
	var head *ListNode
	head = &ListNode{
		Val:  1,
		Next: nil,
	}
	start := head
	for i := 2; i <= num; i++ {
		node := ListNode{
			Val:  i,
			Next: nil,
		}
		start.Next = &node
		start = start.Next
	}
	return head
}

func printList(head ListNode) string {
	var s []int
	start := &head
	for start != nil {
		s = append(s, start.Val)
		start = start.Next
	}
	return fmt.Sprintf("%v", s)
}

func Test_reorderList(t *testing.T) {
	cases := []struct {
		Name     string
		Num      int
		Expected string
	}{
		{"odd", 5, "[1 5 2 4 3]"},
		{"even", 4, "[1 4 2 3]"},
		{"broaden", 1, "[1]"},
		{"long", 10, "[1 10 2 9 3 8 4 7 5 6]"},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			list := createList(c.Num)
			reorderList(list)
			if printList(*list) != c.Expected {
				t.Fatalf("expected %s, but %s got", c.Expected, printList(*list))
			}
		})
	}
}
