package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func newListNode(nums ...int) *ListNode {
	pre := new(ListNode)

	node := pre
	for _, num := range nums {
		node.Next = &ListNode{num, nil}
		node = node.Next
	}
	return pre.Next
}

func (node *ListNode) String() string {
	s := ""
	this := node
	for this != nil {
		s += strconv.Itoa(this.Val)
		if this.Next != nil {
			s += "->"
		}
		this = this.Next
	}
	return s
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	lastNode := head
	cur := head.Next
	for cur != nil {
		if cur.Val != lastNode.Val {
			lastNode.Next = cur
			lastNode = lastNode.Next
		}
		cur = cur.Next
	}
	lastNode.Next = nil
	return head
}

func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return head
}

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	head.Next = deleteDuplicates(head.Next)
	if head.Val == head.Next.Val {
		head.Next = head.Next.Next
	}

	return head
}

func main() {
	node := newListNode(1, 1, 2, 3, 3)
	fmt.Println(deleteDuplicates(node))
}
