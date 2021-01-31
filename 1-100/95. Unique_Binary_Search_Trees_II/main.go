package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return generate(1, n)
}

func generate(start, end int) []*TreeNode {
	res := make([]*TreeNode, 0)

	if start > end {
		res = append(res, nil)
		return res
	}

	for v := start; v <= end; v++ {
		lefts := generate(start, v-1)
		rights := generate(v+1, end)

		for _, left := range lefts {
			for _, right := range rights {
				root := new(TreeNode)
				root.Val = v
				root.Left = left
				root.Right = right
				res = append(res, root)
			}
		}
	}

	return res
}
