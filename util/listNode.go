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

func main() {
	node := newListNode(1, 2, 3, 4)
	fmt.Println(node)

}
