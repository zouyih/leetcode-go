package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxSum := root.Val
	dfsMaxPathSum(root, &maxSum)
	return maxSum
}

func dfsMaxPathSum(node *TreeNode, maxSum *int) int {
	if node == nil {
		return 0
	}
	left := max(0, dfsMaxPathSum(node.Left, maxSum))
	right := max(0, dfsMaxPathSum(node.Right, maxSum))

	*maxSum = max(*maxSum, node.Val+left+right)
	return node.Val + max(left, right)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
