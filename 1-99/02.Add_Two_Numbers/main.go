package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var preHead *ListNode = new(ListNode)
	var node *ListNode = preHead
	var extra, sum int = 0, 0

	for l1 != nil || l2 != nil || extra != 0 {
		sum = extra
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		node.Next = new(ListNode)
		node = node.Next
		node.Val = sum % 10
		extra = sum / 10
	}
	return preHead.Next

}

func main() {
	fmt.Println("vim-go")
}
