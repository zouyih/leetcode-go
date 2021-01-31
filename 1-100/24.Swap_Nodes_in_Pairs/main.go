package main

import "fmt"

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
	s := ""
	for node != nil && node.Next != nil {
		s += fmt.Sprintf("%d -> ", node.Val)
		node = node.Next
	}
	s += fmt.Sprintf("%d", node.Val)

	return s
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	temp := next.Next
	next.Next = head

	head.Next = swapPairs(temp)

	return next
}

func swapPairs2(head *ListNode) *ListNode {
	preHead := new(ListNode)
	preHead.Next = head
	cur := preHead
	for cur.Next != nil && cur.Next.Next != nil {
		first := cur.Next
		second := first.Next

		first.Next = second.Next
		second.Next = first
		cur.Next = second

		cur = first
	}
	return preHead.Next
}

func main() {
	node := makeListNode(1, 2, 3, 4, 5, 6)
	fmt.Println(node)
	fmt.Println(swapPairs(node))

	node = makeListNode(1, 2, 3, 4, 5, 6)
	fmt.Println(swapPairs2(node))
}
