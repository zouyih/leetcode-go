package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func makeListNodes(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	cur := head
	for i := 1; i < len(nums); i++ {
		cur.Next = &ListNode{Val: nums[i]}
		cur = cur.Next
	}
	return head
}

func printListNodes(node *ListNode) {
	if node == nil {
		fmt.Printf("nil")
		return
	}

	for node != nil && node.Next != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Next
	}
	fmt.Printf("%d\n", node.Val)
	return
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fast := head
	slow := head

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	if fast == nil {
		return head.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return head
}

func main() {
	node := makeListNodes([]int{1, 2, 3, 4})
	n := 1
	printListNodes(removeNthFromEnd(node, n))
}
