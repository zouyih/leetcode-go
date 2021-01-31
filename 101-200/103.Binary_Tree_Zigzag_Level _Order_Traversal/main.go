package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	stack1 := []*TreeNode{root}
	stack2 := make([]*TreeNode, 0)

	isInTurn := true
	for len(stack1) > 0 || len(stack2) > 0 {
		nums := make([]int, 0)
		if isInTurn {
			for len(stack1) > 0 {
				node := stack1[len(stack1)-1]
				stack1 = stack1[:len(stack1)-1]
				if node.Left != nil {
					stack2 = append(stack2, node.Left)
				}
				if node.Right != nil {
					stack2 = append(stack2, node.Right)
				}
				nums = append(nums, node.Val)
			}
		} else {
			for len(stack2) > 0 {
				node := stack2[len(stack2)-1]
				stack2 = stack2[:len(stack2)-1]
				if node.Right != nil {
					stack1 = append(stack1, node.Right)
				}
				if node.Left != nil {
					stack1 = append(stack1, node.Left)
				}
				nums = append(nums, node.Val)
			}
		}

		isInTurn = !isInTurn
		res = append(res, nums)
	}
	return res
}

func zigzagLevelOrder1(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	isInTurn := true
	var node *TreeNode
	for len(queue) > 0 {
		count := len(queue)
		nums := make([]int, 0)
		for i := count - 1; i >= 0; i-- {
			node = queue[i]
			if isInTurn {
				if node.Left != nil {
					queue = append(queue, node.Left)
				}
				if node.Right != nil {
					queue = append(queue, node.Right)
				}
			} else {
				if node.Right != nil {
					queue = append(queue, node.Right)
				}
				if node.Left != nil {
					queue = append(queue, node.Left)
				}
			}
			nums = append(nums, node.Val)
		}
		queue = queue[count:]
		isInTurn = !isInTurn
		res = append(res, nums)
	}

	return res
}
