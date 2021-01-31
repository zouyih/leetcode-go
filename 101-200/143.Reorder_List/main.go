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

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	fast = slow.Next
	slow.Next = nil
	fast = reverse(fast)
	slow = head

	for fast != nil {
		slowNext := slow.Next
		slow.Next = fast
		fastNext := fast.Next
		fast.Next = slowNext
		slow = slowNext
		fast = fastNext
	}
}

func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func reorderList1(head *ListNode) {
	list := make([]*ListNode, 0)
	node := head
	for node != nil {
		list = append(list, node)
		node = node.Next
	}
	l, r := 0, len(list)-1
	for l < r {
		list[l].Next = list[r]
		l++
		if l == r {
			break
		}
		list[r].Next = list[l]
		r--
	}
	list[l].Next = nil
}

func main() {
	head := newListNode(1, 2, 3, 4)
	reorderList(head)
	fmt.Println(head)
}
