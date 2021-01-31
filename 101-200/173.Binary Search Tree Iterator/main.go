package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	head  *TreeNode
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	bi := BSTIterator{}
	bi.head = root

	stack := make([]*TreeNode, 0)
	for root != nil {
		stack = append(stack, root)
		root = root.Left
	}
	bi.stack = stack
	return bi
}

func (bi *BSTIterator) Next() int {
	stack := bi.stack
	cur := stack[len(stack)-1]
	stack = stack[:len(stack)-1]

	node := cur.Right
	for node != nil {
		stack = append(stack, node)
		node = node.Left
	}
	bi.stack = stack
	return cur.Val
}

func (bi *BSTIterator) HasNext() bool {
	return len(bi.stack) > 0
}
