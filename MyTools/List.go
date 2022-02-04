package MyTools

import (
	"fmt"
)

//ListNode 单链表节点定义
type ListNode struct {
	Val  int
	Next *ListNode
}

//CreateList 创建一个长度为num的单链表，节点值从0递增
func CreateList(num int) *ListNode {
	if num == 0 {
		return nil
	}
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
	if num == 0 {
		return nil
	}
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
func PrintList(head *ListNode) string {
	var s []int
	start := head
	for start != nil {
		s = append(s, start.Val)
		start = start.Next
	}
	return fmt.Sprintf("%v", s)
}

//CreateListByArray 根据数组创建单链表
func CreateListByArray(num []int) *ListNode {
	if len(num) == 0 {
		return nil
	}
	head := &ListNode{
		Val:  num[0],
		Next: nil,
	}
	ret := head
	for i := 1; i < len(num); i++ {
		head.Next = &ListNode{
			Val:  num[i],
			Next: nil,
		}
		head = head.Next
	}
	return ret
}

//CompareList 比较两个单链表是否相同
func CompareList(l1 *ListNode, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	if !(l1 == nil && l2 == nil) {
		return false
	}
	return true
}
