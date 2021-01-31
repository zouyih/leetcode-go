package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	preHead := new(ListNode)
	preHead.Next = head

	pre := preHead
	cur := head
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
			cur = cur.Next
			continue
		}
		pre = cur
		cur = cur.Next
	}
	return preHead.Next
}
