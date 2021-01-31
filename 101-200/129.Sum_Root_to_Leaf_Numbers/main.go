package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	path := 0
	res := 0
	dfs(root, path, &res)
	return res
}

func dfs(node *TreeNode, path int, res *int) {
	if node == nil {
		return
	}

	path = path*10 + node.Val
	if node.Left == nil && node.Right == nil {
		*res += path
		return
	}

	if node.Left != nil {
		dfs(node.Left, path, res)
	}
	if node.Right != nil {
		dfs(node.Right, path, res)
	}
}
