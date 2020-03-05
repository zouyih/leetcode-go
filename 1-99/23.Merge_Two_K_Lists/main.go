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
	root = &ListNode{Val: l[0]}
	node := root
	for _, v := range l[1:] {
		node.Next = &ListNode{Val: v}
		node = node.Next
	}
	return root
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

func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}
	return helper(lists, 0, n-1)
}

func helper(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}

	if left+1 == right {
		return mergeTwoLists(lists[left], lists[right])
	}

	mid := (left + right) / 2
	l1 := helper(lists, left, mid-1)
	l2 := helper(lists, mid, right)
	return mergeTwoLists(l1, l2)
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	head := new(ListNode)
	node := head
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
	return head.Next
}

func main() {
	l1 := makeListNode([]int{1, 4, 5})
	l2 := makeListNode([]int{1, 3, 4})
	l3 := makeListNode([]int{2, 6})
	lists := []*ListNode{l1, l2, l3}
	fmt.Println(mergeKLists(lists))
}
