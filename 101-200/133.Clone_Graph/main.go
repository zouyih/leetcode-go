package main

type Node struct {
	val       int
	neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	hash := make(map[*Node]*Node)
	return cloneGraphDfs(node, hash)
}

func cloneGraphDfs(node *Node, hash map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}

	if clone, ok := hash[node]; ok {
		return clone
	}

	clone := &Node{val: node.val}
	hash[node] = clone
	clone.neighbors = make([]*Node, len(node.neighbors))
	for i, neighbor := range node.neighbors {
		clone.neighbors[i] = cloneGraphDfs(neighbor, hash)
	}
	return clone
}

func cloneGraphBfs(node *Node) *Node {
	if node == nil {
		return nil
	}

	hash := make(map[*Node]*Node)

	clone := &Node{val: node.val, neighbors: make([]*Node, len(node.neighbors))}
	hash[node] = clone
	queue := []*Node{node}

	for len(queue) != 0 {
		node := queue[0]
		cloneNode := hash[node]
		queue = queue[1:]

		for i, neighbor := range node.neighbors {
			if _, ok := hash[neighbor]; !ok {
				hash[neighbor] = &Node{val: neighbor.val, neighbors: make([]*Node, len(neighbor.neighbors))}
				queue = append(queue, neighbor)
			}
			cloneNode.neighbors[i] = hash[neighbor]
		}
	}
	return clone
}
