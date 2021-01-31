package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	path := make([]int, 0)
	dfs(root, sum, path, &res)
	return res
}

func dfs(node *TreeNode, sum int, path []int, res *[][]int) {
	sum -= node.Val
	path = append(path, node.Val)

	if sum == 0 && node.Left == nil && node.Right == nil {
		*res = append(*res, append([]int{}, path...))
	}

	if node.Left != nil {
		dfs(node.Left, sum, path, res)
	}
	if node.Right != nil {
		dfs(node.Right, sum, path, res)
	}
	path = path[:len(path)-1]
}
