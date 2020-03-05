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

func partition(head *ListNode, x int) *ListNode {
	beforePre := new(ListNode)
	afterPre := new(ListNode)

	before := beforePre
	after := afterPre
	for head != nil {
		if head.Val < x {
			before.Next = head
			before = before.Next
		} else {
			after.Next = head
			after = after.Next
		}
		head = head.Next
	}
	after.Next = nil
	before.Next = afterPre.Next
	return beforePre.Next
}

func main() {
	head := newListNode(1, 4, 3, 2, 5, 2)
	x := 3
	node := partition(head, x)
	fmt.Println(node)
}
