package main

type UF struct {
	count  int
	parent []int
	size   []int
}

func NewUF(n int) *UF {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}

	return &UF{n, parent, size}
}

func (uf *UF) Find(x int) int {
	parent := uf.parent
	for parent[x] != x {
		parent[x] = parent[parent[x]]
		x = parent[x]
	}
	uf.parent = parent
	return x
}

func (uf *UF) Union(p, q int) {
	rootP := uf.Find(p)
	rootQ := uf.Find(q)
	if rootP == rootQ {
		return
	}

	size := uf.size
	if size[rootP] < size[rootQ] {
		uf.parent[rootP] = rootQ
		size[rootQ] += size[rootP]
	} else {
		uf.parent[rootQ] = rootP
		size[rootP] += size[rootQ]
	}
	uf.count--
}

func (uf *UF) isConnected(p, q int) bool {
	rootP := uf.Find(p)
	rootQ := uf.Find(q)

	return rootP == rootQ
}

func (uf *UF) Count() int {
	return uf.count
}
