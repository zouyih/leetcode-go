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

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	headVal := head.Val
	if head.Next != nil && head.Next.Val == headVal {
		for head != nil && head.Val == headVal {
			head = head.Next
		}

		return deleteDuplicates(head)
	}

	head.Next = deleteDuplicates(head.Next)
	return head
}

func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	preHead := &ListNode{}
	preHead.Next = head

	slow := preHead
	fast := head

	for fast != nil {
		if fast.Next != nil && fast.Val == fast.Next.Val {
			val := fast.Val
			for fast != nil && fast.Val == val {
				fast = fast.Next
			}
		} else {
			slow.Next = fast
			slow = slow.Next
			fast = fast.Next
		}
	}

	slow.Next = nil
	return preHead.Next
}

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	preHead := &ListNode{}
	preHead.Next = head

	slow := preHead
	fast := head

	for fast != nil {
		for fast.Next != nil && fast.Val == fast.Next.Val {
			fast = fast.Next
		}
		if slow.Next == fast {
			slow = fast
		} else {
			slow.Next = fast.Next
		}
		fast = fast.Next
	}

	return preHead.Next
}

func main() {
	node := newListNode(1, 2, 3, 3, 4, 4, 5)
	fmt.Println(deleteDuplicates1(node))
}
