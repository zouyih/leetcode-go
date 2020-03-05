package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	n := 0
	node := head
	for node != nil {
		n++
		node = node.Next
	}
	return helper(head, n)
}

func helper(head *ListNode, n int) *TreeNode {
	if n == 0 {
		return nil
	}
	mid := n / 2
	node := head
	for i := 0; i < mid; i++ {
		node = node.Next
	}

	root := &TreeNode{Val: node.Val}
	root.Left = helper(head, mid)
	root.Right = helper(node.Next, n-mid-1)
	return root
}

func sortedListToBST1(head *ListNode) *TreeNode {
	n := 0
	node := head
	for node != nil {
		n++
		node = node.Next
	}
	return helper1(&head, n)
}

func helper1(head **ListNode, n int) *TreeNode {
	if n == 0 {
		return nil
	}
	mid := n / 2

	left := helper1(head, mid)
	root := &TreeNode{Val: (*head).Val}
	root.Left = left
	*head = (*head).Next
	root.Right = helper1(head, n-mid-1)
	return root
}
