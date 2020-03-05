package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftD := minDepth(root.Left)
	rightD := minDepth(root.Right)

	if leftD == 0 {
		return rightD + 1
	}
	if rightD == 0 {
		return leftD + 1
	}
	return min(leftD, rightD) + 1
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
