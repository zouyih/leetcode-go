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

func sortList(head *ListNode) *ListNode {
	n := 0
	node := head
	for node != nil {
		n++
		node = node.Next
	}

	step := 1
	preHead := new(ListNode)
	preHead.Next = head

	for step < n {
		first := preHead.Next
		tail := preHead
		for first != nil {
			second := cutListNode(first, step)
			next := cutListNode(second, step)

			subHead, subTail := mergeListNode(first, second)
			tail.Next = subHead

			first = next
			tail = subTail
		}
		step <<= 1
	}
	return preHead.Next
}

func mergeListNode(first, second *ListNode) (head, tail *ListNode) {
	preHead := new(ListNode)
	node := preHead
	for first != nil && second != nil {
		if first.Val < second.Val {
			node.Next = first
			first = first.Next
		} else {
			node.Next = second
			second = second.Next
		}
		node = node.Next
	}

	for first != nil {
		node.Next = first
		first = first.Next
		node = node.Next
	}

	for second != nil {
		node.Next = second
		second = second.Next
		node = node.Next
	}

	head = preHead.Next
	tail = node
	return
}

func cutListNode(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	pre := head
	for cur != nil && n > 0 {
		pre = cur
		cur = cur.Next
		n--
	}
	pre.Next = nil
	return cur
}


func main() {
	head := newListNode(4, 2, 1, 3)
	fmt.Println(sortList(head))
}
