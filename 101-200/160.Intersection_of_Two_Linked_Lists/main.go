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

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA := listLength(headA)
	lenB := listLength(headB)

	for lenA > lenB {
		headA = headA.Next
		lenA--
	}
	for lenB > lenA {
		headB = headB.Next
		lenB--
	}

	for headA != nil && headB != nil && headA != headB {
		headA = headA.Next
		headB = headB.Next
	}
	return headA
}

func listLength(head *ListNode) int {
	n := 0
	node := head
	for node != nil {
		n++
		node = node.Next
	}
	return n
}

func main() {
	headA := newListNode(1, 2, 3, 7, 8, 9)
	headB := newListNode(4, 5)

	node := headB
	for node.Next != nil {
		node = node.Next
	}

	node.Next = headA.Next.Next.Next
	fmt.Println(headA)
	fmt.Println(headB)
	fmt.Println(getIntersectionNode(headA, headB)) //7->8->9
}
