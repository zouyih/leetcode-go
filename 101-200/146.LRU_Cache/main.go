package main

import "fmt"

type Node struct {
	key  int
	val  int
	pre  *Node
	next *Node
}

func newNode(key, val int) *Node {
	return &Node{key: key, val: val}
}

type DoubleList struct {
	head *Node
	tail *Node
	size int
}

func newDoubleList() *DoubleList {
	head := new(Node)
	tail := new(Node)
	head.next = tail
	tail.pre = head
	return &DoubleList{head: head, tail: tail}
}

func (dl *DoubleList) AddFirst(node *Node) {
	next := dl.head.next
	dl.head.next = node
	node.pre = dl.head
	node.next = next
	next.pre = node
	dl.size++
}

func (dl *DoubleList) Remove(node *Node) {
	pre := node.pre
	next := node.next
	pre.next = next
	next.pre = pre
	dl.size--
}

func (dl *DoubleList) RemoveLast() *Node {
	if dl.tail.pre == dl.head {
		return nil
	}
	last := dl.tail.pre
	dl.Remove(last)
	return last
}

func (dl *DoubleList) Size() int {
	return dl.size
}

type LRUCache struct {
	capacity int
	hash     map[int]*Node
	dl       *DoubleList
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{}
	lru.capacity = capacity
	lru.hash = make(map[int]*Node)
	lru.dl = newDoubleList()
	return lru
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.hash[key]
	if !ok {
		return -1
	}
	this.Put(key, node.val)
	return node.val
}

func (this *LRUCache) Put(key int, value int) {
	dl := this.dl

	if oldNode, ok := this.hash[key]; ok {
		dl.Remove(oldNode)
	} else {
		if dl.Size() == this.capacity {
			lastNode := dl.RemoveLast()
			delete(this.hash, lastNode.key)
		}
	}

	node := newNode(key, value)
	dl.AddFirst(node)
	this.hash[key] = node
}

func main()  {
	cache := Constructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1))      // returns 1
	cache.Put(3, 3)                // evicts key 2
	fmt.Println(cache.Get(2))      // returns -1 (not found)
	cache.Put(4, 4)                // evicts key 1
	fmt.Println(cache.Get(1))      // returns -1 (not found)
	fmt.Println(cache.Get(3))      // returns 3
	fmt.Println(cache.Get(4) )     // returns 4
}