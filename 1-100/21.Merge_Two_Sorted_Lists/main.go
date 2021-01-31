package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func makeListNode(l []int) (root *ListNode) {
	if len(l) == 0 {
		return root
	}
	root = &ListNode{}
	node := root
	for _, v := range l {
		node.Next = &ListNode{Val: v}
		node = node.Next
	}
	return root.Next
}

func printListNode(node *ListNode) {
	for node != nil && node.Next != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Next
	}
	fmt.Printf("%d\n", node.Val)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{}
	node := root
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}
		node = node.Next
	}
	if l1 != nil {
		node.Next = l1
	} else {
		node.Next = l2
	}
	return root.Next
}

func main() {
	l1 := []int{1, 2, 4}
	l2 := []int{1, 3, 4}
	n1 := makeListNode(l1)
	n2 := makeListNode(l2)
	node := mergeTwoLists(n1, n2)
	printListNode(node)
}
