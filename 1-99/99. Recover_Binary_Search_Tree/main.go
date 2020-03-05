package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	var node1, node2, pre *TreeNode
	cur := root
	stack := make([]*TreeNode, 0)

	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pre != nil && cur.Val < pre.Val {
			if node1 == nil {
				node1 = pre
			}
			if node1 != nil {
				node2 = cur
			}
		}
		pre = cur
		cur = cur.Right
	}

	node1.Val, node2.Val = node2.Val, node1.Val
}
