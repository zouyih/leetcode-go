package main

type Node struct {
	val    int
	next   *Node
	random *Node
}

func copyRandomList(head *Node) *Node {
	hash := make(map[*Node]*Node)
	return copyRandomListHelper(head, hash)
}

func copyRandomListHelper(head *Node, hash map[*Node]*Node) *Node {
	if head == nil {
		return nil
	}
	if clone, ok := hash[head]; ok {
		return clone
	}

	clone := &Node{val: head.val}
	clone.next = copyRandomListHelper(head.next, hash)
	clone.random = copyRandomListHelper(head.random, hash)
	hash[head] = clone
	return clone
}
