package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	stack := make([]*TreeNode, 0)
	cur := root
	var pre *TreeNode

	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pre != nil && cur.Val <= pre.Val {
			return false
		}
		pre = cur
		cur = cur.Right
	}
	return true
}

func isValidBST1(root *TreeNode) bool {
	var lastVal int
	hasVal := false
	return isValid(root, &lastVal, &hasVal)
}

func isValid(node *TreeNode, lastVal *int, hasVal *bool) bool {
	if node == nil {
		return true
	}

	if !isValid(node.Left, lastVal, hasVal) {
		return false
	}

	if *hasVal && node.Val <= *lastVal {
		return false
	}

	*lastVal = node.Val
	*hasVal = true
	return isValid(node.Right, lastVal, hasVal)
}
