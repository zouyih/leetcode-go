package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return isSymmetricTwo(root, root)
}

func isSymmetricTwo(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}
	if node1.Val == node2.Val {
		return isSymmetricTwo(node1.Left, node2.Right) && isSymmetricTwo(node1.Right, node2.Left)
	} else {
		return false
	}
}

func isSymmetric1(root *TreeNode) bool {
	stack := []*TreeNode{root, root}

	for len(stack) > 0 {
		n := len(stack)
		node1 := stack[n-1]
		node2 := stack[n-2]
		stack = stack[:n-2]

		if node1 == nil && node2 == nil {
			continue
		}

		if node1 == nil || node2 == nil {
			return false
		}

		if node1.Val != node2.Val {
			return false
		}
		stack = append(stack, node1.Left, node2.Right, node1.Right, node2.Left)
	}
	return true
}
