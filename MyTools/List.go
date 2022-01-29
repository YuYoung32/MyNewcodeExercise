package MyTools

import "fmt"

//ListNode 单链表节点定义
type ListNode struct {
	Val  int
	Next *ListNode
}

//CreateList 创建一个长度为num的单链表，节点值从0递增
func CreateList(num int) *ListNode {
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

//CreateLoopList 创建一个有num个节点的单链表，循环节点为第loop个
func CreateLoopList(num int, loop int) *ListNode {
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

//PrintList 打印单链表，将单链表的各节点值整合成一个切片并打印
func PrintList(head ListNode) string {
	var s []int
	start := &head
	for start != nil {
		s = append(s, start.Val)
		start = start.Next
	}
	return fmt.Sprintf("%v", s)
}
