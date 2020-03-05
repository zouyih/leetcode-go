package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return depth(root) != -1
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := depth(root.Left)
	if leftDepth == -1 {
		return -1
	}

	rightDepth := depth(root.Right)
	if rightDepth == -1 {
		return -1
	}

	d := max(leftDepth, rightDepth)
	if d > leftDepth+1 || d > rightDepth+1 {
		return -1
	}
	return d + 1
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
