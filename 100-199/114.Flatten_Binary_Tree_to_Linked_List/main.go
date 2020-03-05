package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	stack := make([]*TreeNode, 0)
	node := root
	for node != nil {
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			node.Right = node.Left
			node.Left = nil
		} else if len(stack) > 0 {
			node.Right = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		node = node.Right
	}
}

func flatten1(root *TreeNode) {
	if root == nil {
		return
	}

	flatten1(root.Left)
	flatten1(root.Right)
	tmp := root.Right
	root.Right = root.Left
	root.Left = nil

	node := root
	for node.Right != nil {
		node = node.Right
	}
	node.Right = tmp
}
