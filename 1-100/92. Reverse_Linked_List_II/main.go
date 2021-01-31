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

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	preHead := new(ListNode)
	preHead.Next = head
	preNode := preHead
	for i := 1; i < m; i++ {
		preNode = preNode.Next
	}
	tail := preNode.Next

	cur := preNode.Next
	pre := new(ListNode)
	for i := m; i <= n; i++ {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	preNode.Next = pre
	tail.Next = cur
	return preHead.Next
}

func main() {
	head := newListNode(1, 2, 3, 4, 5)
	m := 2
	n := 4
	fmt.Println(reverseBetween(head, m, n))
}
