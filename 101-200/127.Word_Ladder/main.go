package main

import "fmt"

func ladderLength(beginWord string, endWord string, wordList []string) int {
	length := len(beginWord)

	hash := make(map[string][]string)
	for _, word := range wordList {
		for i := 0; i < length; i++ {
			temp := word[:i] + "*" + word[i+1:]
			if _, ok := hash[temp]; !ok {
				hash[temp] = make([]string, 0)
			}
			hash[temp] = append(hash[temp], word)
		}
	}

	visited := make(map[string]bool)
	res := 1
	queue := []string{beginWord}
	for len(queue) != 0 {
		res++
		queLen := len(queue)
		for i := 0; i < queLen; i++ {
			word := queue[0]
			queue = queue[1:]

			for j := 0; j < length; j++ {
				temp := word[:j] + "*" + word[j+1:]
				targets, ok := hash[temp]
				if !ok {
					continue
				}
				for _, target := range targets {
					if target == endWord {
						return res
					}
					if visited[target] {
						continue
					}
					queue = append(queue, target)
					visited[target] = true
				}
			}
		}
	}
	return 0
}

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	fmt.Println(ladderLength(beginWord, endWord, wordList))
}
