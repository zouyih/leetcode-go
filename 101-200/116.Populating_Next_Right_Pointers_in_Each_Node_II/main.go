package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	cur := root
	for cur != nil {
		dummy := new(Node)
		tail := dummy
		for cur != nil {
			if cur.Left != nil {
				tail.Next = cur.Left
				tail = tail.Next
			}
			if cur.Right != nil {
				tail.Next = cur.Right
				tail = tail.Next
			}
			cur = cur.Next
		}
		cur = dummy.Next
	}
	return root
}
