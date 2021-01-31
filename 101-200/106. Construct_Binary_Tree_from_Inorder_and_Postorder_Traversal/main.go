package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	rootVal := postorder[len(postorder)-1]
	root := &TreeNode{Val: rootVal}
	count := 0
	for _, val := range inorder {
		if val != rootVal {
			count++
		} else {
			break
		}
	}

	root.Left = buildTree(inorder[:count], postorder[:count])
	root.Right = buildTree(inorder[count+1:], postorder[count:len(postorder)-1])
	return root
}
