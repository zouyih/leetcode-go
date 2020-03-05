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

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	oldTail := head
	n := 1
	for oldTail.Next != nil {
		n++
		oldTail = oldTail.Next
	}
	oldTail.Next = head

	k = n - k%n - 1
	newTail := head
	for k > 0 {
		newTail = newTail.Next
		k--
	}

	newHead := newTail.Next
	newTail.Next = nil
	return newHead
}

func main() {
	head := newListNode(1, 2, 3, 4, 5)
	k := 2
	fmt.Println(rotateRight(head, k))
}
