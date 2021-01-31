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

func insertionSortList(head *ListNode) *ListNode {
	preHead := new(ListNode)
	preHead.Next = head

	pre := preHead
	cur := head

	for cur != nil {
		next := cur.Next
		for pre.Next != nil && pre.Next.Val < cur.Val {
			pre = pre.Next
		}
		cur.Next = pre.Next
		pre.Next = cur
		if cur.Next == cur {
			cur.Next = nil
		}

		cur = next
		pre = preHead
	}

	return preHead.Next
}

func main() {
	head := newListNode(4, 2, 1, 3)
	fmt.Println(insertionSortList(head))
}
