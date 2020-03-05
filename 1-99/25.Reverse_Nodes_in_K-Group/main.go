package main

import (
	"bytes"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func makeListNode(vals ...int) *ListNode {
	head := new(ListNode)
	node := head
	for _, v := range vals {
		node.Next = &ListNode{Val: v}
		node = node.Next
	}
	return head.Next
}

func (node *ListNode) String() string {
	var out bytes.Buffer
	for node != nil && node.Next != nil {
		s := fmt.Sprintf("%d -> ", node.Val)
		out.WriteString(s)
		node = node.Next
	}
	s := fmt.Sprintf("%d", node.Val)
	out.WriteString(s)

	return out.String()
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	node := head
	for i := 0; i < k; i++ {
		if node == nil {
			return head
		}
		node = node.Next
	}
	newHead := reverse(head, node)
	head.Next = reverseKGroup(node, k)
	return newHead
}

func reverse(a, b *ListNode) *ListNode {
	cur := a
	var pre *ListNode
	for cur != b {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func main() {
	node := makeListNode(1, 2, 3, 4)
	fmt.Println(reverseKGroup(node, 2))
}
