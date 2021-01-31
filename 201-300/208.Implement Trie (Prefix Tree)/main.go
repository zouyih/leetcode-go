package main

import "fmt"

type Trie struct {
	LookUp map[uint8]*Trie
	IsEnd  bool
}

func Constructor() Trie {
	trie := Trie{}
	trie.LookUp = make(map[uint8]*Trie)

	return trie
}

func (trie *Trie) Insert(s string) {
	cur := trie
	for i := range s {
		char := s[i]
		lookUp := cur.LookUp
		if _, ok := lookUp[char]; !ok {
			newTrie := Constructor()
			lookUp[char] = &newTrie
		}
		cur = lookUp[char]
	}
	cur.IsEnd = true
}

func (trie *Trie) Search(s string) bool {
	cur := trie
	for i := range s {
		char := s[i]
		lookUp := cur.LookUp
		if _, ok := lookUp[char]; !ok {
			return false
		}
		cur = lookUp[char]
	}
	return cur.IsEnd
}

func (trie *Trie) StartsWith(s string) bool {
	cur := trie
	for i := range s {
		char := s[i]
		lookUp := cur.LookUp
		if _, ok := lookUp[char]; !ok {
			return false
		}
		cur = lookUp[char]
	}
	return true
}

func main() {
	trie := Constructor()

	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))
	fmt.Println(trie.Search("app"))
	fmt.Println(trie.StartsWith("app"))

	trie.Insert("app")
	fmt.Println(trie.Search("app"))
}
