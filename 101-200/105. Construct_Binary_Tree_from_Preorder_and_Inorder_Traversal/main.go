package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}
	count := 0
	for _, val := range inorder {
		if val != rootVal {
			count++
		} else {
			break
		}
	}

	root.Left = buildTree(preorder[1:1+count], inorder[:count])
	root.Right = buildTree(preorder[1+count:], inorder[1+count:])
	return root
}

func buildTree1(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	cur := root
	stack := []*TreeNode{root}

	pre := 1
	in := 0

	for pre < len(preorder) {
		if cur.Val != inorder[in] {
			cur.Left = &TreeNode{Val: preorder[pre]}
			cur = cur.Left
			stack = append(stack, cur)
			pre++
		} else {
			for len(stack) > 0 && stack[len(stack)-1].Val == inorder[in] {
				cur = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				in++
			}
			cur.Right = &TreeNode{Val: preorder[pre]}
			cur = cur.Right
			stack = append(stack, cur)
			pre++
		}
	}
	return root
}
