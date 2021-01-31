package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	queue := make([]*TreeNode, 0)

	if root == nil {
		return res
	}

	queue = append(queue, root)
	for len(queue) != 0 {
		n := len(queue)
		res = append(res, queue[n-1].Val)

		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return res
}
